package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := []string{""}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	fetch(args...)
}

func fetch(params ...string) {
	const api = "https://www.toptal.com/developers/gitignore/api/"

	param := strings.Join(params, ",")

	uri := api + param

	res, err := http.Get(uri)
	if err != nil {
		log.Println("http request fail")
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
	}

	writeToFile(data)
}

func writeToFile(data []byte) {
	f, err := os.Create(".gitignore")
	if err != nil {
		log.Println("file create fail")
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		log.Println("write file fail")
	}
}
