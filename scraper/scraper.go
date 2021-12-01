package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type AmazonProductInfo struct {
	Title     string   `param:"name" json:"name"`
	ImagesURL []string `param:"imageURL" json:"imageURL"`
	Desc      string   `param:"description" json:"description"`
	Price     string   `param:"price" json:"price"`
	Reviews   string   `param:"totalReviews" json:"totalReviews"`
}

type ProductInfoResponse struct {
	Url     string            `param:"url" json:"url"`
	Product AmazonProductInfo `param:"product" json:"product"`
}

type AmazonProductRequest struct {
	AmazonProductUrl string `param:"producturl" json:"producturl"`
}

//utils
func ParseBody(r *http.Request, body interface{}) (request interface{}) {
	contentType := r.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		json.NewDecoder(r.Body).Decode(&body)
	}
	return body
}

func GetResponse(method string, url string, jsonReq []byte) ([]byte, error) {
	log.Println("Url: ", url)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, _ := client.Do(request)
	return ioutil.ReadAll(response.Body)
}

//----------------------------------

func GetAmazonProductInfo(url string) (p ProductInfoResponse) {
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err.Error())
	})

	p.Url = url
	// title
	c.OnHTML("title", func(e *colly.HTMLElement) {
		p.Product.Title = e.Text
	})

	c.OnHTML("div[id=imageBlock]", func(e *colly.HTMLElement) {
		link := e.ChildAttrs("img", "src")
		for i := 0; i < len(link); i++ {
			if strings.HasSuffix(link[i], "jpg") || strings.HasSuffix(link[i], "png") {
				p.Product.ImagesURL = append(p.Product.ImagesURL, link[i])
			}
		}

	})

	c.OnHTML(`div[id=feature-bullets]`, func(e *colly.HTMLElement) {
		e.ForEach("ul", func(_ int, el *colly.HTMLElement) {
			p.Product.Desc = el.ChildText("li")
		})
	})

	c.OnHTML(`span[id=priceblock_ourprice]`, func(e *colly.HTMLElement) {
		p.Product.Price = e.Text
	})

	c.OnHTML(`span[id=priceblock_dealprice]`, func(e *colly.HTMLElement) {
		p.Product.Price = e.Text
	})

	// reviews
	c.OnHTML(`div[data-hook=total-review-count]`, func(e *colly.HTMLElement) {
		re := regexp.MustCompile("[0-9]+")
		p.Product.Reviews = re.FindString(strings.TrimSpace(e.Text))

	})

	c.Visit(url)

	return
}

func main() {
	setupController()
}

func StoreProductInfo(productDetails ProductInfoResponse) {

	jsonReq, _ := json.Marshal(productDetails)

	_, err := GetResponse("POST", "http://storage-service:8081/save/product/amazon", jsonReq)
	if err != nil {
		panic(err)
	}
}

func GetProductInfo(w http.ResponseWriter, r *http.Request) {
	body := AmazonProductRequest{}
	request := ParseBody(r, body)

	mResult := request.(map[string]interface{})
	url := mResult["producturl"]

	productDetails := GetAmazonProductInfo(url.(string))
	jsonReq, err := json.Marshal(productDetails)
	if err != nil {
		panic(err)
	}

	StoreProductInfo(productDetails)

	io.WriteString(w, string(jsonReq))

}

func setupController() {
	http.HandleFunc("/product/amazon", GetProductInfo)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
