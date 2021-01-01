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

	val, err := client.Create("burn_test", "1234", "vehis83986@cocyo.com", 2000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", val)

	// val, err := client.Burn("dv2ypby01n8qygj2qzpluhbq0yb88mf")
	// if err != nil {
	// 	fmt.Println("ERROR:", err)
	// }
	// fmt.Printf("B: %+v\n", val)
}
