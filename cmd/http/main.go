package main

import (
	"log"
	"sagara-msib-test/internal/boot"
)

func main() {
	if err := boot.BajuInventoryHTTP(); err != nil {
		log.Printf("HTTP failed to boot server due to : %v\n", err.Error())
	}
}
