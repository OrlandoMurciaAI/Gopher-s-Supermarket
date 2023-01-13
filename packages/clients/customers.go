package client

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Customer structure
type Customer struct {
	Name    string
	Id      int
	Age     int
	Address string
}

var data []byte

const filePathNames = "files/names.txt"
const filePathEntrance = "files/Entrance.csv"

func init() {

	var err error
	//Reading the name data"
	data, err = ioutil.ReadFile(filePathNames)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
}

func randomName(i int64) string {
	// Convert the byte slice to a string
	names := string(data)

	// Split the string into a slice of names
	namesSlice := strings.Split(names, "\n")

	//Seed the random number generator
	rand.Seed(time.Now().UnixNano() + i)

	// Pick a random index within the range of the slice
	index := rand.Intn(len(namesSlice))

	// Return the selected name
	return strings.TrimSpace(namesSlice[index])
}

// Function that generates a slice of customers
func GenerateCustomers(quantity int) []Customer {
	var customers []Customer
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < quantity; i++ {
		customers = append(customers, Customer{
			Name:    randomName(int64(i)),
			Id:      r.Intn(999999999) + 1,
			Age:     r.Intn(100) + 1,
			Address: "Address" + strconv.Itoa(i+1),
		})
	}
	return customers
}

func GreetingCustomers(customer []Customer) {
	for _, value := range customer {
		fmt.Println("Welcome to the Gopher Supermarket: ", value.Name)
		entrance(value)
	}

}

func entrance(customer Customer) {
	dt := time.Now()
	// Opening data for the entrance
	file, err := os.OpenFile(filePathEntrance, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// Handle the error
		fmt.Println("Error reading file", err)
	}
	//Create the csv write
	writer := csv.NewWriter(file)
	writer.Write([]string{strconv.Itoa(customer.Id), strconv.Itoa(customer.Age), customer.Address, dt.Format("01-02-2006 15:04:05"), customer.Name})
	writer.Flush()
	defer file.Close()
}
