package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/fatih/color"
)

const baseURL = "https://storage.googleapis.com/"

func checkBucket(bucket string) {
	url := baseURL + bucket
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer requisição para:", url)
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var responseBody string
	for scanner.Scan() {
		responseBody += scanner.Text()
	}

	if strings.Contains(responseBody, "<Code>AccessDenied</Code>") {
		color.Yellow("[Access Denied] %s", bucket)
	} else if strings.Contains(responseBody, "<Code>NoSuchBucket</Code>") {
		color.Red("[No Such Bucket] %s", bucket)
	} else {
		color.Green("[VULNERABLE] %s", bucket)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <wordlist>")
		os.Exit(1)
	}

	wordlistPath := os.Args[1]
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("Erro ao abrir wordlist:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bucket := strings.TrimSpace(scanner.Text())
		if bucket != "" {
			checkBucket(bucket)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler wordlist:", err)
	}
}
