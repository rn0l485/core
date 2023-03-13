package web

import (
	"fmt"
	"time"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/rn0l485/core/utility"
	"github.com/rn0l485/core/setting"
)

type LogParams struct {
	// TimeStamp shows the time after the server returns a response.
	TimeStamp time.Time `json:"TimeStamp" bson:"timestamp"`
	// StatusCode is HTTP response code.
	StatusCode int `json:"StatusCode" bson:"statuscode"`
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration `json:"Latency" bson:"latency"`
	// ClientIP equals Context's ClientIP method.
	ClientIP string `json:"ClientIP" bson:"clientip"`
	// Method is the HTTP method given to the request.
	Method string `json:"Method" bson:"method"`
	// Path is a path the client requests.
	Path string `json:"Path" bson:"path`
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage string `json:"ErrorMessage" bson:"errormessage"`

	// BodySize is the size of the Response Body
	BodySize int `json:"BodySize" bson:"bodysize"`
	// Keys are the keys set on the request's context.
	Keys map[string]interface{} `json:"Keys" bson:"keys"`
	// contains filtered or unexported fields	
}

func Logger(msgC chan string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// before request

		c.Next()

		// after request

		param := LogParams{
			Keys:    c.Keys,
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		// json making 
		if b, err := json.Marshal(param); err != nil {
			fmt.Println(err)
		} else {
			utility.ChannelTimeOut(msgC, string(b), time.Second*10)
		}
	}
}

func PageNotFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"Msg": "page-not-found",
		"StatusCode":404,
	})
}

func Alive(c *gin.Context) {
	c.JSON( http.StatusOK, gin.H{
		"Msg":"ok",
		"StatusCode":200,
	})
}

func Done(c *gin.Context, data ...interface{}) {
	if len(data) == 0 {
		Alive(c)
	} else if len(data) == 1 {
		c.JSON( http.StatusOK, gin.H{
			"Msg":"ok",
			"StatusCode":200,
			"Data": data[0],
		})
	} else {
		c.JSON( http.StatusOK, gin.H{
			"Msg":"ok",
			"StatusCode":200,
			"Data": data,
		})
	}
}

func Error(c *gin.Context, errCode interface{}, err error, netCode ...int) {
	var httpCode int 
	if len(netCode) == 0 {
		httpCode = http.StatusNotFound
	} else {
		httpCode = netCode[0]
	}

	if err != nil{
		c.AbortWithStatusJSON( httpCode, gin.H{
			"Msg": errCode,
			"StatusCode": httpCode,
			"Description": err.Error(),
		})
	} else {
		c.AbortWithStatusJSON( httpCode, gin.H{
			"Msg": errCode,
			"StatusCode": httpCode,
		})		
	}
}

func GinNewRouterInit() (*gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}

func GinRouterInit() (*gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}

func GinRouterWithLogInit() (*gin.Engine) {
	return gin.Default()
}

func GinSession(R *gin.Engine, path, domain, secret string, maxAge int) error {
	if path == "" || domain == "" || secret == "" || maxAge < 60 {
		return setting.Err_NoData
	} 

	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{
		Path: path,
		Domain: domain,
		MaxAge: maxAge,
		Secure: false,
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
	})
	R.Use(sessions.Sessions("client", store))

	return nil
}






/*
func GinJWT(R *gin.Engine, RealmName, secret, IdentityKey string) error {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:			RealmName,
		Key:			[]byte(secret),
		Timeout:		time.Hour,
		MaxRefresh:		time.Hour,
		IdentityKey: 	identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
		  if v, ok := data.(*User); ok {
		    return jwt.MapClaims{
		      identityKey: v.UserName,
		    }
		  }
		  return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
		  claims := jwt.ExtractClaims(c)
		  return &User{
		    UserName: claims[identityKey].(string),
		  }
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
		  var loginVals login
		  if err := c.ShouldBind(&loginVals); err != nil {
		    return "", jwt.ErrMissingLoginValues
		  }
		  userID := loginVals.Username
		  password := loginVals.Password

		  if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
		    return &User{
		      UserName:  userID,
		      LastName:  "Bo-Yi",
		      FirstName: "Wu",
		    }, nil
		  }

		  return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
		  if v, ok := data.(*User); ok && v.UserName == "admin" {
		    return true
		  }

		  return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
		  c.JSON(code, gin.H{
		    "code":    code,
		    "message": message,
		  })
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})	
}
*/