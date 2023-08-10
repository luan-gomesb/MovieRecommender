package main

import (
	"errors"
	"fmt"
	"luan-gomesb/MovieRecommender/openaiapi"
	"net/http"
	"strings"
)

func getToken(r *http.Request) (string, error) {
	var gptToken string
	if r.Method != "POST" {
		return "", errors.New("Request method needs to be POST")
	}
	if r.Header.Get("Authorization") == "" {
		return "", errors.New("OpenAI needs to be sended as Bearer token")
	}

	userToken := r.Header.Get("Authorization")

	if apitokenArr := strings.Split(userToken, " "); len(apitokenArr) == 2 {
		gptToken = apitokenArr[1]
	}
	return gptToken, nil
}

// "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"
func hangleGet(w http.ResponseWriter, r *http.Request) {
	userToken, e := getToken(r)
	if e != nil {
		fmt.Fprintf(w, e.Error())
		return
	}
	response, e := openaiapi.Openaiapi(userToken)

	if e != nil {
		fmt.Fprintf(w, e.Error())
		return
	}
	fmt.Fprintf(w, response)
}

func main() {
	http.HandleFunc("/", hangleGet)
	http.ListenAndServe(":8080", nil)
}
