package src

import (
	"fmt"
	"net/http"
)

func shagowing() error {
	var (
		client *http.Client
		err    error
	)

	if true {
		client, err = createTracingClient()
	} else {
		client, err = createTracingClient()
	}

	if err != nil {
		return err
	}

	client.Timeout = 10

	fmt.Println("client:", client)
	return nil
}

func createTracingClient() (*http.Client, error) {
	transport := &http.Transport{
		// configure transport options if needed
	}
	return &http.Client{
		Transport: transport,
	}, nil
}
