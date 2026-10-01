package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)


func main () {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the domain name you want to scan")
	fmt.Printf("This will provide you the following information:")
	fmt.Println("MX records,", "SPF Records,", "DMARC Records")
	fmt.Println("------------------------------------------------")
	

	for scanner.Scan() {
		domain := scanner.Text()
		domain = strings.TrimSpace(domain)
		fmt.Println("Scanning domain:", domain,"...")
		checkDomain(domain)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading input:", err)
	}

}


func checkDomain (domain string) {

	var hasMX, hasSPF, hasDMARC bool
	var spfRecords, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("Error looking up MX records:", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Println("Error looking up TXT records:", err)
	}

	if len(txtRecords) > 0 {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecords = record
			}
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	
	if err != nil {
		fmt.Println("Error looking up DMARC records:", err)
	}

	if len(dmarcRecords) > 0 {
		hasDMARC = true
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				dmarcRecord = record
			}
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecords, hasDMARC, dmarcRecord)


}