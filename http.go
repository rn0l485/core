package core

import (
	"errors"
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
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
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}
}

func StartHttpServer( addr string, handler http.Handler ) error {
	return HttpServer(addr, handler).ListenAndServe()
}




func NewReq(method, url string, data ...interface{}) (*http.Request, error) {
	switch method {
	case "POST":
		if len(data) != 0 {
			return NewJsonPost(url, data[0])
		} else {
			return NewJsonPost(url, make(map[string]string{}))
		}
	case "FORM":
		if len(data) != 0 {
			return NewFormPost(url, data[0])
		} else {
			return NewJsonPost(url, make(map[string]string{}))
		}
	case "GET":
		return NewGet(url)
	default:
		return nil, errors.New("method assign error")
	}
}

func NewFormPost(url string, data interface{}) (*http.Request, error) {
	d, ok := data.(url.Values)
	if !ok {
		return nil, errors.New("NewFormPost.assertion: input data type error.")
	}

	req, err := http.NewRequest( "POST", url, strings.NewReader(d.Encode()))
	if err != nil {
		return nil, errors.New("NewFormPost.http.NewRequest:"+err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func NewJsonPost(url string, data interface{}) ( *http.Request, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("NewJsonPost.json.Marshal:"+err.Error())
	}

	req, err := http.NewRequest( "POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.New("NewJsonPost.http.NewRequest:"+err.Error())
	}

	return req, nil 
}

func NewGet( url string ) ( *http.Request, error ) {
	if req, err := http.NewRequest( "GET", url, nil); err != nil {
		return nil, errors.New("NewGet.http.NewRequest."+err.Error())
	} else {
		return req, nil 
	}
}