// https://zetcode.com/golang/getpostrequest/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// For this case values may be nil too
	values := map[string]string{
		"name":       "John Doe",
		"occupation": "gardener",
	}

	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal("json", err)
	}

	resp, err := http.Post("http://localhost:8081/articles", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal("Response:", err)
	}

	fmt.Println(resp.Status)
}
