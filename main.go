/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "log"

func main() {
	server, err := InitializeServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	server.Start()
}
