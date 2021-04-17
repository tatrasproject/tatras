package main

import (
	"fmt"
	// "log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

const url = "https://argocd.site.com/api/v1/applications"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Type %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	fmt.Printf("Response: %v \n", content)

	documents := documentFromJson(content)
	for _, document := range documents {
		fmt.Println(document.Name)
	}

}

func documentFromJson(content string) []Document {
	documents := make([]Document, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var document Document
	for decoder.More() {
		err := decoder.Decode(&document)
		if err != nil {
			panic(err)
		}
		documents = append(documents, document)
	}
	return documents

}

type Document struct {
	Name, File string
}