package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	setupController()
}

func GetIndianTimeStampNow() (time_now time.Time) {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	now := time.Now().In(loc)
	return now
}

func WriteAmazonProductInfoToFile(w http.ResponseWriter, r *http.Request) {

	date_time := "\"createdAt\":" + "\"" + GetIndianTimeStampNow().String() + "\","

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("Error during reading body: %v", err)
	}

	//avoid duplicate product data saving in file
	re := regexp.MustCompile("\"(url)\"[:](\"([^\"\"]+)\")")
	re2 := regexp.MustCompile("\"(name)\"[:](\"([^\"\"]+)\")")
	match := re.FindStringSubmatch(string(body))
	match2 := re2.FindStringSubmatch(string(body))
	b, err := ioutil.ReadFile("output.txt")
	if err != nil {
		panic(err)
	}
	s := string(b)
	if strings.Contains(s, match[0]) || strings.Contains(s, match2[0]) {
		return
	}

	//add timestamp
	p := string(body)
	index := 1
	data_with_timestamp := p[:index] + date_time + p[index:]

	//open file to write amazon product details
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	if _, err = f.WriteString(data_with_timestamp + "\n"); err != nil {
		panic(err)
	}

	defer f.Close()

}

func GetFileData(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("output.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	b, err := ioutil.ReadAll(file)

	delimited := strings.Replace(string(b), "\n", ",", -1)
	final := "[" + strings.TrimSuffix(delimited, ",") + "]"

	io.WriteString(w, final)

}

func setupController() {
	http.HandleFunc("/save/product/amazon", WriteAmazonProductInfoToFile)
	http.HandleFunc("/data", GetFileData)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
