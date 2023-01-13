package main

import (
	"math/rand"
	"time"

	cs "GOPHER-S-SUPERMAKET/packages/clients"
)

func main() {
	// Generate a slice of 10 customers
	ch1 := make(chan []cs.Customer)

	// Customer generation
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			ch1 <- cs.GenerateCustomers(rand.Intn(10))
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	// This loop configures the simulation time
	for i := 0; i < 1; i++ {
		customerBatch := <-ch1
		cs.GreetingCustomers(customerBatch)
	}
}
