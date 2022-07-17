package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Isi struct {
	Judul string `json:"judul"`
	Slug  string `json:"slug"`
}

type Isinya []Isi

func main() {
	resp, err := http.Get("https://github.zenia.my.id/terbaru.json")
	cek(err)

	content, errornya := ioutil.ReadAll(resp.Body)
	cek(errornya)

	var isinya Isinya

	err = json.Unmarshal(content, &isinya)
	cek(err)

	// ini hasilnya
	// fmt.Println(isinya)

	readmeFile, err := ioutil.ReadFile("README.md")
	cek(err)
	readme := string(readmeFile)

}
