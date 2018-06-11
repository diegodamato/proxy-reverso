package main

import (
	"proxy-reverso/internal/app"
	"proxy-reverso/internal/pkg/model"

	"github.com/jinzhu/configor"
)

func init() {
	configor.Load(&model.Config, "./configs/config.yml")
}

func main() {
	app.Bootstrap()
}
