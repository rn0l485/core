package web

import (
	"errors"
	"net/http"
	"net/url"
	"time"
	"encoding/json"
	"bytes"
	"strings"
	"golang.org/x/sync/errgroup"
)

func HttpServer( addr string, handler http.Handler ) *http.Server {
	return &http.Server{
		Addr:			addr,
		Handler:		handler,
		ReadTimeout:	10 * time.Second,
		WriteTimeout:	10 * time.Second,
	}
}

func StartMultiHttpServer(ServerList ...*http.Server) error {
	var g errgroup.Group
	for i:=0; i<len(ServerList); i++ {
		g.Go(func() error {
			err := ServerList[i].ListenAndServe()
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	} 

	return nil
}

func StartHttpServer( addr string, handler http.Handler ) error {
	return HttpServer(addr, handler).ListenAndServe()
}

func NewReq(method, target string, data ...interface{}) (*http.Request, error) {
	switch method {
	case "POST":
		if _, ok := data[0].(map[string]interface{}); !ok {
			return nil, errors.New("NewReq.assertion: input data type error")
		}

		if len(data) != 0 {
			return NewJsonPost( target, data[0])
		} else {
			return NewJsonPost( target, make(map[string]string, 0))
		}
	case "FORM":
		//if _, ok := data[0].( url.Values); !ok {
		//	return nil, errors.New("NewReq.assertion: input data type error.")
		//}
		if len(data) != 0 {
			return NewFormPost( target, data[0])
		} else {
			return NewJsonPost( target, url.Values{})
		}
	case "GET":
		return NewGet( target)
	default:
		return nil, errors.New("method assign error")
	}
}

func NewFormPost( target string, data interface{}) (*http.Request, error) {
	d, ok := data.( url.Values)
	if !ok {
		return nil, errors.New("NewFormPost.assertion: input data type error.")
	}

	req, err := http.NewRequest( "POST",  target, strings.NewReader(d.Encode()))
	if err != nil {
		return nil, errors.New("NewFormPost.http.NewRequest:"+err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form- targetencoded")

	return req, nil
}

func NewJsonPost( target string, data interface{}) ( *http.Request, error) {
	if _, ok := data.(map[string]interface{}); !ok {
		return nil, errors.New("NewJsonPost.assertion: input data type error")
	}
	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("NewJsonPost.json.Marshal:"+err.Error())
	}

	req, err := http.NewRequest( "POST",  target, bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.New("NewJsonPost.http.NewRequest:"+err.Error())
	}

	return req, nil 
}

func NewGet(  target string ) ( *http.Request, error ) {
	if req, err := http.NewRequest( "GET", target, nil); err != nil {
		return nil, errors.New("NewGet.http.NewRequest."+err.Error())
	} else {
		return req, nil 
	}
}