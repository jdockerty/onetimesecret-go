package main

import (
	"fmt"
	"os"

	"github.com/jdockerty/onetimesecret-go/ots"
)

func main() {
	var c ots.Client

	client := c.New(os.Getenv("OTS_EMAIL"), os.Getenv("OTS_KEY"))
	client.Status()

	val, err := client.Create("super_secret_text", "password1234", "vehis83986@cocyo.com", 60)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
	
}
