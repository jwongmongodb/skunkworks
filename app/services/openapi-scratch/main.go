package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// set OpenAI API key as environment variable
	apiKey, err := ioutil.ReadFile("api.key")
	if err != nil {
		fmt.Printf("Error reading API key file: %v\n", err)
		return
	}

	authorization := fmt.Sprintf("Bearer %s", apiKey)
	println(authorization)

	// set request parameters
	url := "https://api.openai.com/v1/images/generations"
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", apiKey),
	}
	requestBody := map[string]interface{}{
		"model":           "image-alpha-001",
		"prompt":          "a beautiful sunset over the mountains",
		"num_images":      1,
		"size":            "512x512",
		"response_format": "url",
	}

	// send HTTP POST request to OpenAI API
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("Error marshaling request body: %v\n", err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// parse response body
	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		fmt.Printf("Error decoding response body: %v\n", err)
		return
	}
	fmt.Printf("Generated image URL: %v\n", responseBody["data"].([]interface{})[0].(map[string]interface{})["url"])
}
