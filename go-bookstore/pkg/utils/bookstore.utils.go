package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Unmarshalling JSON

func ParseBody(r *http.Request, data interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		fmt.Println("Without Byte Body - ", body)
		fmt.Println("Byte Body - ", []byte(body))
		if err := json.Unmarshal([]byte(body), data); err != nil {
			return
		}
	}
}
