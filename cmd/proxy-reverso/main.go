package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"proxy-reverso/internal/pkg/model"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/configor"
)

func main() {
	ProxyReverse()
}

func init() {
	config := model.Config{}
	configor.Load(&config, "./configs/config.yml")
}

func HostReverseProxy(target url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}
	return &httputil.ReverseProxy{Director: director}
}

func ProxyReverse() {

	proxy := HostReverseProxy(
		url.URL{
			Scheme: "http",
			Host:   "localhost:3001",
			Path:   "/jogos",
		})

	router := mux.NewRouter()
	router.HandleFunc("/teste", proxy.ServeHTTP)

	app := negroni.New()
	app.UseHandler(router)
	app.Run(":8083")
}
