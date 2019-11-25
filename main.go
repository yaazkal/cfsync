package main

import (
	"fmt"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/rdegges/go-ipify"
)

func main() {
	// Check requirements
	fmt.Println("Checking requirements...")
	if os.Getenv("CF_API_KEY") == "" || os.Getenv("CF_MAIL") == "" || os.Getenv("CF_ZONE_NAME") == "" || os.Getenv("CF_A_RECORD") == "" {
		fmt.Println("Please set all the system environment variables")
		os.Exit(0)
	}

	// Set variables
	CF_API_KEY := os.Getenv("CF_API_KEY")
	CF_MAIL := os.Getenv("CF_MAIL")
	CF_ZONE_NAME := os.Getenv("CF_ZONE_NAME")
	CF_A_RECORD := os.Getenv("CF_A_RECORD")
	targetRecord := CF_A_RECORD + "." + CF_ZONE_NAME

	// Get public IP
	fmt.Println("Gathering the public IP...")
	ip, err := ipify.GetIp()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The IP to use is %s\n", ip)

	fmt.Printf("The zone name to use is %s\n", CF_ZONE_NAME)
	// Create the cloudflare object to be used
	cfapi, err := cloudflare.New(CF_API_KEY, CF_MAIL)
	if err != nil {
		panic(err)
	}

	// Fetch the zone ID for zone host
	zoneID, err := cfapi.ZoneIDByName(CF_ZONE_NAME)
	if err != nil {
		panic(err)
	}

	// Check if the record exist by name and type
	recordQuery := cloudflare.DNSRecord{Name: targetRecord, Type: "A"}
	records, err := cfapi.DNSRecords(zoneID, recordQuery)
	if err != nil {
		panic(err)
	}

	// Decide if the record needs to be updated or created.
	if len(records) == 0 {
		fmt.Println("No records found. The record " + targetRecord + " will be created.")
		recordQuery.Content = ip
		_, err := cfapi.CreateDNSRecord(zoneID, recordQuery)
		if err != nil {
			panic(err)
		}
		fmt.Println("Record created succesfully.")
	} else {
		fmt.Println("Record found. The record " + targetRecord + " will be updated.")
		recordID := records[0].ID
		recordQuery.Content = ip
		err := cfapi.UpdateDNSRecord(zoneID, recordID, recordQuery)
		if err != nil {
			panic(err)
		}
		fmt.Println("Record updated succesfully.")
	}

}
