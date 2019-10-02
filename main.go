package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("----Service Checker----")
	fmt.Println("-----------------------")
	filename := os.Args[1] //let it explode just in case..
	file, e := os.Open(filename)
	defer file.Close()
	if e != nil {
		panic(e)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsed := parseEnv(scanner.Text())
		if parsed != "" {
			checkServiceIsAlive(parsed)
		}
	}
}

func checkServiceIsAlive(endpoint string) {
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("Argh! Error calling %s\n", endpoint)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Printf("%s OK\n", endpoint)
	} else {
		fmt.Printf("Argh! %s - %d\n", endpoint, resp.StatusCode)
	}
}

func parseEnv(env string) string {
	if env == "" {
		return ""
	}
	s := strings.Split(env, "=")
	url := s[1]
	if strings.HasPrefix(url, "http://") {
		return url
	}
	return ""
}
