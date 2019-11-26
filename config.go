package main

type Config struct {
	API_key string
	Email   string
	Zones   map[string][]Record
}
