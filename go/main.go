package main

import (
	"conversion/shp_to_json"
	"io/ioutil"
	"log"
	"strings"
)

/*
my thoughs on this code:
we should accept one or more files from the frontend.
we either return directly one object or multiple objects as a simple json array, depending on the input
the output should simply be a list of objects and the objects directly containing their geomtry
geomtry only contains, type and all vertecies
*/
func main() {
	dirname := "../backend/files"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for index, file := range files {
		if index > 3 {
			break
		}

		filename := file.Name()
		if !file.IsDir() && strings.HasSuffix(filename, ".shp") {
			absPath := dirname + "/" + filename
			var output = shp_to_json.Convert(absPath)
			log.Default().Println(output)
		} else {
			log.Default().Println(filename)
		}
	}
}
