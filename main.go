package main

import (
	"log"
)

func main() {
	// Check requirements and read configuration file
	config := readConfiguration()

	// Configuration information summary
	log.Printf("Found %d zones in configuration file\n", len(config.Zones))

	// Start sync
	log.Println("Starting syncronization process...")
	ip := publicIP()
	syncronize(config, ip)
}
