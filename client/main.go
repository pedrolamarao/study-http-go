package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func erase(client *http.Client) error {
	key := os.Args[2]
	request, err := http.NewRequest(http.MethodDelete, "http://127.0.0.1:8080/"+key, nil)
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func get(client *http.Client) error {
	key := os.Args[2]
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/"+key, nil)
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	log.Print(string(bytes))
	return nil
}

func put(client *http.Client) error {
	key := os.Args[2]
	value := os.Args[3]
	request, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:8080/"+key, strings.NewReader(value))
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("usage: tool (get|put|erase) key [value]")
	}

	var cmd func(*http.Client) error
	switch os.Args[1] {
	case "delete":
		cmd = erase
		break
	case "get":
		cmd = get
		break
	case "put":
		cmd = put
		break
	default:
		log.Fatal("usage: tool (get|put|delete) key [value]")
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	err := cmd(client)
	if err != nil {
		log.Fatal(err)
	}
}
