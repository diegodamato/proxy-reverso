package model

type Location struct {
	Origin      string
	Destination string
}

type Config struct {
	Address []Location
}
