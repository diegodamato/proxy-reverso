package reverseProxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"proxy-reverso/internal/pkg/model"
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

	for key := range AddressMap {
		proxy := hostReverseProxy(
			url.URL{
				Scheme: "http",
				Host:   AddressMap[key].Destination,
				Path:   AddressMap[key].DestinationPath,
			})

		http.HandleFunc(AddressMap[key].OriginPath, proxy.ServeHTTP)
	}

	log.Println("Server running on port 8082")
	log.Fatal(http.ListenAndServeTLS(":8082", "./certificates/server.crt", "./certificates/server.key", nil))

}
