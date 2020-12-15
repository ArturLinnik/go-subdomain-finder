// All the subdomains lists are from https://github.com/rbsec/dnscan

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// Asks for the number of subdomains you want to try out
	fmt.Println("Numbers of subdomains to try out (100, 1000 or 10000): ")
	var numberSubdomains int
	fmt.Scanln(&numberSubdomains)

	// Uses a list depending on the users choice
	var subdomainList string

	if numberSubdomains == 100 {
		subdomainList = "subdomains-100.txt"
	} else if numberSubdomains == 1000 {
		subdomainList = "subdomains-1000.txt"
	} else if numberSubdomains == 10000 {
		subdomainList = "subdomains-10000.txt"
	} else {
		fmt.Println("Error.")
		os.Exit(1)
	}

	// Gets every line of the text document
	file, err := os.Open(subdomainList)

	if err != nil {
		log.Fatalf("Failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	// Asks for the domain name
	var domain string
	fmt.Printf("Enter your domain: ")
	fmt.Scanln(&domain)

	// Do requests to the url using every subdomain of the list and prints put those which exists
	for _, subdomain := range text {
		url := "http://" + subdomain + "." + domain
		_, err := http.Get(url)
		if err == nil {
			fmt.Println(url)
		}
	}
}
