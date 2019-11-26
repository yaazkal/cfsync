package main

import (
	"io/ioutil"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/rdegges/go-ipify"
	"gopkg.in/yaml.v2"
)

func readConfiguration() Config {
	log.Println("Getting configuration...")
	var config Config

	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func publicIP() string {
	log.Println("Gathering public IP...")
	ip, err := ipify.GetIp()
	if err != nil {
		panic(err)
	}
	log.Printf("The IP to use is %s\n", ip)
	return ip
}

func syncronize(config Config, ip string) {
	// Create the cloudflare object to be used
	cfapi, err := cloudflare.New(config.API_key, config.Email)
	if err != nil {
		panic(err)
	}

	for zone, records := range config.Zones {
		// Fetch the zone ID
		zoneID, err := cfapi.ZoneIDByName(zone)
		if err != nil {
			panic(err)
		}

		for _, record := range records {
			// Check if the record exist by name and type
			targetRecord := record.Name + "." + zone
			recordQuery := cloudflare.DNSRecord{Name: targetRecord, Type: record.Type}
			resultQuery, err := cfapi.DNSRecords(zoneID, recordQuery)
			if err != nil {
				panic(err)
			}
			// Decide if the record needs to be updated or created.
			if len(resultQuery) == 0 {
				log.Println("The record " + targetRecord + " will be created.")
				recordQuery.Content = ip
				_, err := cfapi.CreateDNSRecord(zoneID, recordQuery)
				if err != nil {
					panic(err)
				}
				log.Println("Record created succesfully.")
			} else {
				if resultQuery[0].Content != ip {
					log.Println("The record " + targetRecord + " will be updated.")
					recordID := resultQuery[0].ID
					recordQuery.Content = ip
					err := cfapi.UpdateDNSRecord(zoneID, recordID, recordQuery)
					if err != nil {
						panic(err)
					}
					log.Println("Record updated succesfully.")
				} else {
					log.Println("The record " + targetRecord + " doesn't need to be updated.")
				}
			}
		}
	}
}
