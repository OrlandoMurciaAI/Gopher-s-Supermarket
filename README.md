# Gopher's Supermarket

This project is a simple simulation of a supermarket using the Go programming language. It is designed to help you practice and improve your Go skills by providing examples of various concepts and techniques.

## Features
Generates a list of customers with random names, ID, age, and address.
Writes the entrance of the customers in a CSV file.
Prints a greeting message for each customer.

## Getting Started
To get started with the project, you will need to have Go installed on your computer. You can download the latest version of Go from the official website (https://golang.org/dl/).

Once you have Go installed, you can clone the repository by running the following command:
```
git clone https://github.com/OrlandoMurciaAI/Gopher-s-Supermarket.git
```

## Running the simulation
To run the simulation, you can navigate to the project's root directory and run the following command:

```
go run main.go
```

This will generate a list of customers and write their entrance in a CSV file and print the greeting message for each customer.

## Project Structure
The project is structured as follows:

* main.go: The entry point of the program. It contains the main function that runs the simulation.
* customer.go: It has the customer structure and the functions to generate the list of customers and print the greeting message.
* files: It has the result files called "Entrance.csv" and names.txt which it is used for picking  a random name 