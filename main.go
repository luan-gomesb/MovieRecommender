package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Choice struct {
	Text          string `json:"text"`
	Index         int    `json:"index"`
	Logprobs      string `json:"logprobs"`
	Finish_reason string `json:"finish_reason"`
}
type Usage struct {
	Prompt_tokens     int `json:"prompt_tokens"`
	Completion_tokens int `json:"completion_tokens"`
	Total_tokens      int `json:"total_tokens"`
}

type ChatResponse struct {
	Object  string   `json:"object"`
	Id      string   `json:"id"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

func main2() {
	askChatgpt()
	// textjson();
}

func askChatgpt() {

	apiToken := os.Getenv("OPENAI_API_KEY")
	data := bytes.NewBuffer([]byte(`{
	"model": "text-davinci-003",
	"prompt": "album mais vendido do ACDC",
	"temperature": 0,
	"max_tokens": 500,
	"top_p": 1,
	"frequency_penalty": 0.0,
	"presence_penalty": 0.0
	}`))

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add the authorization header to the request
	req.Header.Add("Authorization", "Bearer "+apiToken)
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf("Body : %s",string(body))
	var response ChatResponse
	json.Unmarshal(body, &response)
	fmt.Println(response.Choices[0].Text)
}
