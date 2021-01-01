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


	val, err := client.Generate("1234", "vehis83986@cocyo.com", 20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", val)

}
