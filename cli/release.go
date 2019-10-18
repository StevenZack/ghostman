package main

import (
	"io/ioutil"
	"log"
	"os"
)

var res = make(map[string]string)

func main() {
	registerRes()
	release()
}

func registerRes() {
	resList := []string{
		"web/index.html",
		"web/index.tis",
	}

	for _, path := range resList {
		res[path] = readFile(path)
	}

	
}

func release() {

}

func readFile(f string) string {
	file, e := os.OpenFile(f, os.O_RDONLY, 0644)
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	b, e := ioutil.ReadAll(file)
	if e != nil {
		log.Fatal(e)
	}
	return string(b)
}
