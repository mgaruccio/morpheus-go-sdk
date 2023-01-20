package main

import (
	"fmt"
	"log"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func main() {
	client := morpheus.NewClient("https://yourmorpheus.com")
	client.SetUsernameAndPassword("username", "password")
	resp, err := client.Login()
	if err != nil {
		fmt.Println("LOGIN ERROR: ", err)
	}
	fmt.Println("LOGIN RESPONSE:", resp)

	// List load balancer types
	req := &morpheus.Request{}
	loadBalancerTypeResponse, err := client.ListLoadBalancerTypes(req)
	if err != nil {
		log.Fatal(err)
	}
	result := loadBalancerTypeResponse.Result.(*morpheus.ListLoadBalancerTypesResult)
	log.Println(&result.LoadBalancerTypes)
}
