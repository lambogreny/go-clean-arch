package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://172.30.0.77:8882/api/v1/technicalViability"

	payload := strings.NewReader("{\n\t\"lat\": \"0\",\n\t\"lng\": \"0\",\n\t\"city\": \"augusto\",\n\t\"state\": \"ckpmcwioecwre\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
