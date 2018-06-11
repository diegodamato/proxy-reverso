package model

type Location struct {
	OriginPath      string
	Destination     string
	DestinationPath string
}

type Configuracao struct {
	Address []Location
}

var Config = Configuracao{}
