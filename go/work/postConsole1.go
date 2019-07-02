package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://10.0.10.100/console"

	payload := strings.NewReader("{\n\t\"user\": \"root\",\n    \"password\": \"calvin\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "068df98f-44dc-491b-a2e9-db28c9cea13c")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("post log failed: %s", http.StatusText(res.StatusCode))
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
