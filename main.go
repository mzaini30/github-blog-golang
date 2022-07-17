package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
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

	isi := ""
	for _, x := range isinya {
		isi += "- [" + x.Judul + "](https://github.zenia.my.id/tulisan/" + x.Slug + ")\n"
	}
	isi = "## Blog\n\n" + isi
	// println(isi)

	readmeFile, err := ioutil.ReadFile("README.md")
	cek(err)
	readme := string(readmeFile)

	sampelRegex := regexp.MustCompile("(<!-- blog start -->)(\n.*?)(<!-- blog end -->)")
	hasil := sampelRegex.ReplaceAllString(readme, "$1\n"+isi+"$3")
	// println(hasil)

	hasilFile := []byte(hasil)

	error := ioutil.WriteFile("README.md", hasilFile, 0644)
	cek(error)
}
