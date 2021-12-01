package utils

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func GetResponse(method string, url string, jsonReq []byte) ([]byte, error) {
	log.Println("Url: ", url)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, _ := client.Do(request)
	return ioutil.ReadAll(response.Body)
}