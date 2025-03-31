package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

func main() {
	var ctx context.Context
	method := flag.String("method", "GET", "HTTP method (GET or POST)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("Usage: client -method [GET|POST] serverURL resourcePath")
	}
	serverURL, resourcePath := args[0], args[1]

	fullURL, err := buildURL(serverURL, resourcePath)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := sendRequest(ctx, *method, fullURL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(body))
}

func buildURL(serverURL, resourcePath string) (string, error) {
	if !strings.HasPrefix(serverURL, "http://") && !strings.HasPrefix(serverURL, "https://") {
		serverURL = "http://" + serverURL
	}
	u, err := url.Parse(serverURL)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, resourcePath)
	return u.String(), nil
}

func sendRequest(ctx context.Context, method, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	return client.Do(req)
}
