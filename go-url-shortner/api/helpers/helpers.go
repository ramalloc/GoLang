package helpers

import (
	"os"
	"strings"
)

// func EnforceHTTP(url string) string {
// 	if url[:4] != "http"{
// 		return "http://" + url
// 	}
// 	return url
// }

// Optimized 
func EnforceHTTP(url string) string {
	if !strings.HasPrefix(url, "http") {
		return "http://" + url
	}
	return url
}

// func RemoveDomainError(url string) bool{
// 	if url == os.Getenv("DOMAIN"){
// 		return false
// 	}

// 	newURL := strings.Replace(url, "http://", "", 1)
// 	newURL = strings.Replace(newURL, "https://", "", 1)
// 	newURL = strings.Replace(newURL, "www.", "", 1)
// 	newURL = strings.Split(newURL, "/")[0]

// 	if newURL == os.Getenv("DOMAIN"){
// 		return false
// 	}

// 	return true
// }

// Optimized
// Checks if the domain is restricted
func IsUrlAndDomainSame(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	// Trimming/Removing http:// from URL if it is present
	newURL := strings.TrimPrefix(url, "http://")
	// Trimming/Removing https:// from URL if it is present
	newURL = strings.TrimPrefix(newURL, "https://")
	// Trimming/Removing www. from URL if it is present
	newURL = strings.TrimPrefix(newURL, "www.")

	// Spliting URL after removal of https/http/www. 
	// Splitting on the bases "/" if url like this short.com/something then url[0] will be short.com
	newURL = strings.Split(newURL, "/")[0]

	// Returning True if url and domain are not same
	return newURL != os.Getenv("DOMAIN")
}