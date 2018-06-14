package main

import (
	"flag"
	"log"
	"os"
	"proxy-reverso/internal/app"
	"proxy-reverso/internal/pkg/model"

	"github.com/jinzhu/configor"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

const productionEnv = "PRODUCTION"

func loadFlags() (string, string) {
	configLocation := flag.String("config-location", "./configs/config.yml", "a string")
	environment := flag.String("globo-environment", "DEVELOPMENT", "a string")
	flag.Parse()

	return *configLocation, *environment
}

func init() {

	configLocation, environmentFlag := loadFlags()
	configor.Load(&model.Config, configLocation)

	if environmentFlag == productionEnv || os.Getenv("GLOBO-ENVIRONMENT") == productionEnv {
		log.SetOutput(&lumberjack.Logger{
			Filename:   "/var/log/proxy-reverso.log",
			MaxSize:    1,
			MaxBackups: 14,
			MaxAge:     28,
		})
	} else {
		log.SetOutput(os.Stdout)
	}

	log.Println("Configurações carregadas!")
}

func main() {
	app.Bootstrap()
}
