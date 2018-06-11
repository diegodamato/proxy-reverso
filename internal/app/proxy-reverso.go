package app

import (
	"log"
	"proxy-reverso/internal/pkg/reverseProxy"
)

//Bootstrap - responsible for starting the application
func Bootstrap() {
	log.Println("Iniciando a aplicação...")
	reverseProxy.ReverseProxy()
}
