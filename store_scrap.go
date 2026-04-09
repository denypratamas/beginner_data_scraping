package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Product represents the scraped product data
type Product struct {
	Name  string
	Price string
	Link  string
}

// getStoreData fetches product data from a given URL
func getStoreData(url string) ([]Product, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %v", err)
	}

	var products []Product

	doc.Find("li.product").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Find("h2.woocommerce-loop-product__title").Text())
		
		price := strings.TrimSpace(s.Find("span.woocommerce-Price-amount.amount").Text())
		
		link, _ := s.Find("a").Attr("href")

		product := Product{
			Name:  name,
			Price: price,
			Link:  link,
		}
		products = append(products, product)
	})

	return products, nil
}

// getAllPages scrapes data from all pages up to maxPages
func getAllPages(url string, maxPages int) []Product {
	var allData []Product

	for pageNum := 1; pageNum <= maxPages; pageNum++ {
		pagedURL := fmt.Sprintf("%s/page/%d/", url, pageNum)
		fmt.Printf("Scraping page %d\n", pageNum)

		pageData, err := getStoreData(pagedURL)
		if err != nil {
			log.Printf("Error scraping page %d: %v", pageNum, err)
			continue
		}

		fmt.Printf("Products found: %d\n", len(pageData))

		if len(pageData) == 0 {
			fmt.Println("No more products found, stopping.")
			break
		}

		allData = append(allData, pageData...)
	}

	return allData
}

// saveToCSV saves the product data to a CSV file
func saveToCSV(data []Product, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"name", "price", "link"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing header: %v", err)
	}

	// Write data rows
	for _, product := range data {
		row := []string{product.Name, product.Price, product.Link}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error writing row: %v", err)
		}
	}

	fmt.Printf("Data scraped and saved to %s\n", filename)
	return nil
}

func main() {
	url := "https://scrapeme.live/shop/"
	maxPages := 48

	allData := getAllPages(url, maxPages)
	
	err := saveToCSV(allData, "store_data.csv")
	if err != nil {
		log.Fatalf("Error saving to CSV: %v", err)
	}
}