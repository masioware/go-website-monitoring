package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
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
}

func testWebsite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "has been successfully loaded!")
	} else {
		fmt.Println("Site:", site, "is having problems.", "-> Status Code:", response.StatusCode)
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
