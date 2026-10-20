package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	url := "https://www.confirmtkt.com/train-schedule/09465"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Browser-like headers
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
			"AppleWebKit/537.36 (KHTML, like Gecko) "+
			"Chrome/130.0.0.0 Safari/537.36")
	req.Header.Set("Accept",
		"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed: %s", resp.Status)
	}

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	rows := doc.Find("tbody").First().Find("tr")

	if rows.Length() > 0 {
		rows.Slice(0, rows.Length()-1).Each(func(i int, row *goquery.Selection) {
			cols := row.Find("td")
			if cols.Length() > 0 {
				cols.Slice(0, cols.Length()-2).Each(func(j int, cell *goquery.Selection) {
					if j == 1 {
						fmt.Printf("  Column (%d,%d): %s\n", i, j, strings.TrimSpace(strings.Split(strings.TrimSpace(cell.Text()), "-")[1]))
					} else {
						fmt.Printf("  Column (%d,%d): %s\n", i, j, strings.TrimSpace(cell.Text()))
					}
				})
			}
		})
	}
}
