package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("webscraper.io"),
	)

	var items []item
	c.OnHTML(".product-wrapper", func(h *colly.HTMLElement) {
		item := item{
			Name: h.ChildText(".title"),
			Price: h.ChildText(".price"),
			ImgUrl: h.ChildAttr("img", "src"),
		}

		items = append(items, item)
	})

	err := c.Visit("https://webscraper.io/test-sites/e-commerce/allinone")
	if err != nil {
		fmt.Println("Error visiting website:", err)
	}

	content, err := json.Marshal(items)

	if err != nil {
		panic(err)
	}

	os.WriteFile("products.json", content, 0644)
}