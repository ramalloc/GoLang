package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

type Parser interface {
	getSEOData(res *http.Response) (SeoData, error)
}

type DefaultParser struct {
}

func (df DefaultParser) getSEOData(res *http.Response) (SeoData, error) {
	// We will send res to go-query
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		fmt.Println("Error while getting doc from goquery")
		return SeoData{}, err
	}
	result := SeoData{}
	result.URL = res.Request.URL.String()
	result.StatusCode = res.StatusCode
	result.Title = doc.Find("title").First().Text()
	result.H1 = doc.Find("h1").First().Text()
	result.MetaDescription, _ = doc.Find("meta[name^=description]").Attr("content")
	return result, nil
}

// userAgents
var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func RandomUserAgents() string {
	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)
	randNum := localRand.Intn(len(userAgents))
	return userAgents[randNum]
}

func makeRequest(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error while making request for url : ", url)
		return nil, err
	}
	req.Header.Set("User-Agent", RandomUserAgents())
	res, clientErr := client.Do(req)
	if clientErr != nil {
		return nil, clientErr
	}
	return res, nil
}

// We will check that the particular page is sitemap or not by seeing that page has .xml file or not
func isSitemap(urls []string) ([]string, []string) {
	sitemapFiles := []string{}
	pages := []string{}

	for _, page := range urls {
		foundSitemap := strings.Contains(page, "xml")
		if foundSitemap {
			fmt.Println("Found Sitemap : ", page)
			sitemapFiles = append(sitemapFiles, page)
		} else {
			pages = append(pages, page)
		}
	}
	return sitemapFiles, pages
}

func scrapURLs(extractedURLs []string, parser Parser, concurrency int) []SeoData {
	tokens := make(chan struct{}, concurrency)
	var n int
	n++
	WorkList := make(chan []string)
	results := []SeoData{}

	go func() { WorkList <- extractedURLs }()

	for ; n > 0; n-- {
		list := <-WorkList
		for _, url := range list {
			if url != "" {
				n++
				go func(url string, token chan struct{}) {
					log.Printf("Requesting URL %s", url)
					res, err := ScrapPage(url, tokens, parser)
					if err != nil {
						log.Printf("Encountered Error at URL : %s", url)
					} else {
						results = append(results, res)
					}
					WorkList <- []string{}
				}(url, tokens)
			}
		}
	}

	return results
}

func ScrapPage(url string, tokens chan struct{}, parser Parser) (SeoData, error) {
	crawledPageData, err := crawlPage(url, tokens)
	if err != nil {
		return SeoData{}, err
	}
	res, err := parser.getSEOData(crawledPageData)
	if err != nil {
		return SeoData{}, err
	}
	return res, nil
}

func extractUrls(res *http.Response) ([]string, error) {
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}
	results := []string{}

	// selecting location from doc
	sel := doc.Find("loc")
	// looping on node as got nodes from doc.Find
	for i := range sel.Nodes {
		loc := sel.Eq(i)
		result := loc.Text()
		results = append(results, result)
	}
	return results, nil
}

func extractSiteMapURLs(startURL string) []string {
	WorkList := make(chan []string)
	toCrawl := []string{}

	go func() { WorkList <- []string{startURL} }()

	var n int
	n++
	for ; n > 0; n-- {
		list := <-WorkList
		for _, link := range list {
			n++
			go func(link string) {
				response, err := makeRequest(link)
				if err != nil {
					fmt.Printf("Error retrieving URL : %s", link)
				}
				urls, extractErr := extractUrls(response)
				if extractErr != nil {
					fmt.Printf("Error extracting document from response, URL : %s", link)
				}
				sitemapFiles, pages := isSitemap(urls)
				if sitemapFiles != nil {
					WorkList <- sitemapFiles
				}

				toCrawl = append(toCrawl, pages...)
			}(link)

		}
	}
	// It returning slice to strings/urls which needs to crawl
	return toCrawl
}

// Perform get request on urls using go-query then convert the response into json
func crawlPage(urls string, tokens chan struct{}) (*http.Response, error) {
	tokens <- struct{}{}
	res, err := makeRequest(urls)
	<-tokens
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ScrapSitemap(url string, parser Parser, concurrency int) []SeoData {
	// Extract All urls from site in result and pass into scrapURLs
	results := extractSiteMapURLs(url)
	res := scrapURLs(results, parser, concurrency)
	return res
}

func main() {
	p := DefaultParser{}

	// Passing sitemap URL in the function and saving in results
	results := ScrapSitemap("https://www.quicksprout.com/sitemap.xml", p, 10)
	for _, res := range results {
		fmt.Println(res)
	}

}
