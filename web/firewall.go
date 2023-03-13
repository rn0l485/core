package web


import (
    "time"
    "net"
	"net/http/httputil"
    "crypto/x509"
    "crypto/tls"
    "io/ioutil"
)



func InitFirewall(Port []string, Forward func(*gin.Context), Wall func() gin.HandlerFunc) {
    router := GinNewRouterInit()
    router.Use(Wall())//中间件，起拦截器的作用
    router.Any("/*action", Forward)//所有请求都会经过Forward函数转发

    for i:=0; i<len(Port); i++ {
    	router.Run(":"+Port[i])
    }
}

func CreateForward(c *gin.Context, Host string, IsHttps bool, CAPath string) func(*gin.Context) {
    return func(*gin.Context) {
        targetHost := &httputil.TargetHost{
            Host: Host,
            IsHttps: IsHttps,
            CAPath: CAPath,
        }
        HostReverseProxy(c.Writer, c.Request, targetHost)        
    }
}

func HostReverseProxy(w http.ResponseWriter, req *http.Request, targetHost *httputil.TargetHost) {
    host := ""
    if targetHost.IsHttps {
        host = host + "https://"
    } else {
        host = host + "http://"
    }

    remote, err := url.Parse(host + targetHost.Host)
    if err != nil {
        log.Errorf("url parse error: %s", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    proxy := httputil.NewSingleHostReverseProxy(remote)

    if targetHost.IsHttps {
        tls, err := GetVerTLSConfig(targetHost.CAPath)
        if err != nil {
            log.Errorf("https crt error: %s", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        var pTransport http.RoundTripper = &http.Transport{
            Dial: func(netw, addr string) (net.Conn, error) {
                c, err := net.DialTimeout(netw, addr, time.Second*5)
                if err != nil {
                    return nil, err
                }
                return c, nil
            },
            ResponseHeaderTimeout: time.Second * 5,
            TLSClientConfig: tls,
        }
        proxy.Transport = pTransport
    }

    proxy.ServeHTTP(w, req)
}


func GetVerTLSConfig(CAPath string) (*tls.Config, error) {
    var TlsConfig *tls.Config

    caData, err := ioutil.ReadFile(CAPath)
    if err != nil {
        log.Errorf("read wechat ca fail", err)
        return nil, err
    }
    pool := x509.NewCertPool()
    pool.AppendCertsFromPEM(caData)

    TlsConfig = &tls.Config{
        RootCAs: pool,
    }
    return TlsConfig, nil
}

func MiddleWare() gin.HandlerFunc {
    return func(c *gin.Context) {
        account := c.Request.Header.Get("ename")//从请求头中获取ename字段
        if account == "" {
            c.JSON(http.StatusOK, httputil.Response{
                Code:   400002,
                Message: "用户未登录",
            })
            c.Abort()
            return
        }
            fmt.Println("before middleware")
            c.Set("request", "clinet_request")
            c.Next()
            fmt.Println("before middleware")
    }
}