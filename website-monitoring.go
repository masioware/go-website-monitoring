package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const nMonitoring = 3
const timeMonitoring = 10 * time.Minute

func main() {
	showIntro()

	for {
		showMenu()

		command := inputCommand()
		fmt.Println()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("I don't know this command...")
			os.Exit(-1)
		}

		fmt.Println()
	}
}

func showIntro() {
	version := 1.0

	fmt.Println("Starting website monitoring program ")
	fmt.Println("This program is in version", version)
	fmt.Println()
}

func showMenu() {
	fmt.Println("-> 1 Start Monitoring")
	fmt.Println("-> 2 Show Logs")
	fmt.Println("-> 0 Exit")
}

func inputCommand() int {
	var command int

	fmt.Print("Command: ")
	fmt.Scan(&command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := readWebsitesFromFile()
	for i := 0; i <= nMonitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testWebsite(site)
		}

		time.Sleep(timeMonitoring)
	}
}

func showLogs() {
	fmt.Println("Displaying Logs...")
	fmt.Println()

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(string(file))
}

func testWebsite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "has been successfully loaded!")
		saveLogs(site, true)
	} else {
		fmt.Println("Site:", site, "is having problems.", "-> Status Code:", response.StatusCode)
		saveLogs(site, false)
	}
}

func readWebsitesFromFile() []string {
	fileName := "websites.txt"
	result := []string{}

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(fileName, " -> Error:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		line = strings.TrimSpace(line)
		result = append(result, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return result
}

func saveLogs(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}

	currentTime := time.Now().Format("02/01/2006 15:04:05")
	line := currentTime + " - " + site + " - ONLINE: " + strconv.FormatBool(status)

	file.WriteString(line + "\n")
	file.Close()
}
