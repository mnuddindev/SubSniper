package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	SubFinder()
	SubLister()
	one := FileReader("subfinder.txt")
	two := FileReader("sublister.txt")
	FileWriter(one, two)
	SeparateByStatus()
	CNAMEchecker()
}

func FileReader(name string) []string {
	fileRead, err := os.Open(name)
	if err != nil {
		log.Fatal("Site Lists Not Found")
	}

	fileScanner := bufio.NewScanner(fileRead)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	fileRead.Close()

	return lines
}

func FileWriter(a []string, b []string) []string {
	fmt.Println("Finding for Duplicate sites")
	fmt.Println("Removing Duplicates Items from List")
	check := make(map[string]int)
	appenD := append(a, b...)
	result := make([]string, 0)
	for _, value := range appenD {
		check[value] = 1
	}
	for letter, _ := range check {
		result = append(result, letter)
	}

	all, err := os.Create("all.txt")
	if err != nil {
		log.Fatal("Error Creating New File")
	} else {
		fmt.Println("Creating New File all.txt")
	}

	for _, writeEach := range result {
		write, err := all.WriteString(writeEach + "\n")
		if err != nil {
			log.Fatal("Error while Removing Duplicate Items")
			all.Close()
		} else {
			fmt.Println("Duplicates Item Removed Writing into all.txt")
		}

		_ = write
	}

	return result
}

func SubFinder() {
	com := [5]string{"subfinder", "-d", "", "-o", "subfinder.txt"}
	com[2] = os.Args[1]
	cmd := exec.Command(com[0], com[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Failed to Get Subdomains From SubFinder")
	}
	fmt.Println("SubFinder> " + string(output))
}

func SubLister() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	com := [6]string{"python", "sublist3r.py", "-d", "", "-o", path + "/sublister.txt"}
	com[3] = os.Args[1]
	cmd := exec.Command(com[0], com[1:]...)
	cmd.Dir = path + "/Sublist3r"
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Failed to Get Subdomains From SubList3r")
	}
	fmt.Println("SubList3r> " + string(output))
}

func SeparateByStatus() {
	lines := FileReader("all.txt")

	notFound, err := os.Create("404.txt")
	if err != nil {
		log.Fatal(err)
	}
	found, err := os.Create("200.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, eachline := range lines {
		fmt.Println("Reading Each line")
		url := "http://" + eachline
		response, err := http.Get(url)
		fmt.Println(url)
		if err != nil {
			fmt.Println("No Hostname Register")
			continue
		}

		statusCode := response.StatusCode
		if statusCode == 404 {
			w, err := notFound.WriteString(eachline + "\n")
			if err != nil {
				log.Fatal(err)
				notFound.Close()
			}

			_ = w
			fmt.Println("Listing site With status Code 404 into 404.txt")
		}

		if statusCode == 200 || statusCode == 300 {
			w, err := found.WriteString(eachline + "\n")
			if err != nil {
				log.Fatal(err)
				found.Close()
			}

			_ = w
			fmt.Println("Listing site with Status Code 200 and 301 in 200.txt")
		}
	}
}

func CNAMEchecker() {
	fmt.Println("Checking CNAME")
	cmd := exec.Command("apt", "install", "dnsutils")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("DNS Tool installation Failed")
	}
	cmd2 := exec.Command("python", "cname.py")
	output, err := cmd2.Output()
	if err != nil {
		fmt.Println("CNAME Checking Failed")
	}

	fmt.Println(string(output))

	fmt.Println("Done!! Check Possible.txt for Final Takeover possible Sites")
	fmt.Println("Thanks")
}
