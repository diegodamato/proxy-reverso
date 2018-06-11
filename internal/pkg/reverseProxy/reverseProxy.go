package reverseProxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"proxy-reverso/internal/pkg/model"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func hostReverseProxy(target url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}
	return &httputil.ReverseProxy{Director: director}
}

func mapAddress() map[string]model.Location {
	var addressMap = make(map[string]model.Location)
	for i := 0; i < len(model.Config.Address); i++ {
		addressMap[model.Config.Address[i].OriginPath] = model.Config.Address[i]
	}
	return addressMap
}

// ReverseProxy - responsible for uploading the reverse proxy server
func ReverseProxy() {
	var AddressMap = make(map[string]model.Location)
	AddressMap = mapAddress()
	router := mux.NewRouter()

	for key := range AddressMap {
		proxy := hostReverseProxy(
			url.URL{
				Scheme: "http",
				Host:   AddressMap[key].Destination,
				Path:   AddressMap[key].DestinationPath,
			})

		router.HandleFunc(AddressMap[key].OriginPath, proxy.ServeHTTP)
	}

	app := negroni.New()
	app.UseHandler(router)
	app.Run(":8082")
}
