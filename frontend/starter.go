package main

import (
	"huihuCrawler02/frontend/controller"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	templateFiles = []string{
		"huihuCrawler02/frontend/view/template.html",
		"view/template.html",
	}
)

func main() {
	pathPrefix := "huihuCrawler02/frontend/view"
	http.Handle("/", http.FileServer(http.Dir(pathPrefix)))

	for _, filename := range templateFiles {
		if PathExist(filename) {
			pathPrefix = getPath(filename)
			http.Handle("/search", controller.CreateSearchResultHandler(filename))
			break
		}
	}


	err := http.ListenAndServe(":8088", nil)

	if err != nil {
		panic(err)
	}

}

func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func getPath(filename string) string {
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
