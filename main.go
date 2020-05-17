package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"gopkg.in/gookit/color.v1"
)

func main() {
	Info()
	SubFinder()
	SubLister()
	one := FileReader("subfinder.txt")
	two := FileReader("sublister.txt")
	FileWriter(one, two)
	SeparateByStatus()
	CNAMEchecker()
	FileMover()
}

func Info() {
	logo := `
   _____       _     _____       _                 
  / ____|     | |   / ____|     (_)                
 | (___  _   _| |__| (___  _ __  _ _ __   ___ _ __ 
  \___ \| | | | '_ \\___ \| '_ \| | '_ \ / _ \ '__|
  ____) | |_| | |_) |___) | | | | | |_) |  __/ |   
 |_____/ \__,_|_.__/_____/|_| |_|_| .__/ \___|_|   
                                  | |              
                                  |_| Inad Islam
`

	color.Green.Println(logo)

}

func FileReader(name string) []string {
	fileRead, err := os.Open(name)
	if err != nil {
		color.Red.Println("[🔴] Site Lists Not Found")
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
	color.Yellow.Println("[!] Finding for Duplicate sites")
	color.Yellow.Println("[-] Removing Duplicates Items from List")
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
		color.Red.Println("[🔴] Error Creating New File")
	} else {
		color.Green.Println("[🖍️] Creating New File all.txt")
	}

	for _, writeEach := range result {
		write, err := all.WriteString(writeEach + "\n")
		if err != nil {
			color.Red.Println("[🔴] Error while Removing Duplicate Items")
			all.Close()
		} else {
			color.Yellow.Println("[-] Duplicates Item Removed Writing into all.txt")
		}

		_ = write
	}

	return result
}

func SubFinder() {
	com := [5]string{"subfinder", "-d", "", "-o", "subfinder.txt"}
	com[2] = os.Args[1]
	cmd := exec.Command(com[0], com[1:]...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		color.Red.Println("[🔴] Failed to Get Subdomains From SubFinder")
	}
	color.Green.Println("[🔍] Collecting Site By SubFinder")
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
	_, err = cmd.CombinedOutput()
	if err != nil {
		color.Red.Println("[🔴] Failed to Get Subdomains From SubList3r")
	}
	color.Green.Println("[🔍] Collecting Site By Sublist3r")
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
		fmt.Println("[📖] Reading Each line")
		url := "http://" + eachline
                fmt.Println(url)
		response, err := http.Get(url)
		if err != nil {
			color.Warn.Tips("No Hostname Register")
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
			color.Yellow.Println("[🖍️] Listing site With status Code 404 into 404.txt")
		}

		if statusCode == 200 || statusCode == 300 {
			w, err := found.WriteString(eachline + "\n")
			if err != nil {
				log.Fatal(err)
				found.Close()
			}

			_ = w
			color.Green.Println("[🖍️] Listing site with Status Code 200 and 301 in 200.txt")
		}
	}
}

func CNAMEchecker() {
	fmt.Println("[🔍] Checking CNAME")
	cmd := exec.Command("apt", "install", "dnsutils")
	_, err := cmd.Output()
	if err != nil {
		color.Red.Println("[🔴] DNS Tool installation Failed")
	}
	cmd2 := exec.Command("python", "cname.py")
	output, err := cmd2.Output()
	if err != nil {
		color.Red.Println("[🔴] CNAME Checking Failed")
	}

	fmt.Println(string(output))

	color.Green.Println("[📣] Done!! Check Possible.txt for Final Takeover possible Sites")
	color.Green.Println("Thanks")
}

func FileMover() {
	folder := os.Args[1]
	name := [6]string{"200.txt", "404.txt", "all.txt", "subfinder.txt", "sublister.txt", "possible.txt"}
	move := exec.Command("mkdir", folder)
	_, err := move.Output()
	if err != nil {
		color.Warn.Tips("[🔴] Failed to Move Files")
	}
	cmd := exec.Command("mv", name[0], name[1], name[2], name[3], name[4], name[5], "./"+folder)
	_, err = cmd.Output()
	if err != nil {
		color.Warn.Tips("[🔴] Failed to Move Files")
	} else {
		color.Green.Println("[📣] All Files Moves to Domai Folder")
	}
}
