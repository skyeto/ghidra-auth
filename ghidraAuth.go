package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading username from stdio")
		os.Exit(1)
	}
	username = strings.TrimSpace(username)

	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading password from stdio")
		os.Exit(1)
	}
	password = strings.TrimSpace(password)

	apiUrl := os.Args[4]
	reqBody, err := json.Marshal(map[string]interface{} {
		"username": username,
		"password": password,
	})
	if err != nil {
		fmt.Printf("Failed to json marshal auth request")
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Printf("Failed to create auth request")
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send auth request, %v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body")
		os.Exit(1)
	}
	respText := string(respBody)

	if respText != "ok" {
		fmt.Printf("NOT AUTHORIZED %s, REASON %s", username, respText)
		os.Exit(1)
	} else {
		fmt.Printf("Authorized %s, %s", username, respText)
	}

	repositoriesAdmin := os.Args[2]
	repository := os.Args[3]
	filename := filename()
	content := fmt.Sprintf("-add %s\n-grant %s +r %s\n-grant %s +w %s\n", username, username, repository, username, repository)

	err = os.WriteFile(fmt.Sprintf("%sadm%s.cmd", repositoriesAdmin, filename), []byte(content), 0644)
	time.Sleep(100 * time.Millisecond) // wait .1s for ghidra to catch up

	os.Exit(0)
}

var letters = []rune("0123456789")
func filename() string {
	buf := make([]rune, 10)
	for i := 0; i < 10; i++ {
		buf[i] = letters[rand.Intn(len(letters))]
	}
	return string(buf)
}
