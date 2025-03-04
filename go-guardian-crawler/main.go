package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

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

func GetRequet(targetURL string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", targetURL, nil)
	req.Header.Set("User-Agent", RandomUserAgents())

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}

}

func DiscoverLinks(response *http.Response, baseURL string) []string {
	// Return empty slice if response is nil
	if response == nil {
		return []string{}
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return []string{}
	}

	foundURLs := []string{}

	// Find all <a> elements and extract href attributes
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if res, exists := s.Attr("href"); exists {
			foundURLs = append(foundURLs, res)
		}
	})

	return foundURLs
}
func ChackRealtive(href string, baseURL string) string {
	if strings.HasPrefix(href, "/") {
		fmt.Printf("%s%s", baseURL, href)
	} else {
		return href
	}
	return href
}

func ResolveRealtiveLinks(link string, baseURl string) (bool, string) {
	resultLink := ChackRealtive(link, baseURl)
	baseParse, _ := url.Parse(baseURl)
	resultParse, _ := url.Parse(resultLink)

	if baseParse != nil && resultParse != nil {
		if baseParse.Host == resultParse.Host {
			return true, resultLink
		} else {
			return false, ""
		}
	}
	return false, ""
}

// Gave five threads at one time to a channel
var tokens = make(chan struct{}, 5)

func Crawl(targetURL string, baseURL string) []string {
	fmt.Println(targetURL)
	// As we use multiple channels to make different request it can overload the site therefore we use Semaphore
	// Semaphore use to control number of channels being used. Below we are using tokens
	tokens <- struct{}{}
	response, _ := GetRequet(targetURL)
	<-tokens
	links := DiscoverLinks(response, baseURL)
	foundLinks := []string{}
	for _, link := range links {
		ok, correctLink := ResolveRealtiveLinks(link, baseURL)
		if ok {
			if correctLink != "" || correctLink != "" {
				foundLinks = append(foundLinks, correctLink)
			}
		}
	}

	return foundLinks
}

func main() {
	WorkList := make(chan []string)
	var n int
	n++
	baseDomain := "https://theguardian.com"
	go func() { WorkList <- []string{baseDomain} }()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <- WorkList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string, baseURL string) {
					foundLinks := Crawl(link, baseDomain)
					if foundLinks != nil {
						WorkList <- foundLinks
					}
				}(link, baseDomain)
			}
		}
	}
}
