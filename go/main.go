package main

import (
	"fmt"
	"os"
	"shp_to_geojson"
)

func main() {
	bytes, err := os.ReadFile("./A1_NODE.shp")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(shp_to_geojson.ParseFromBytes(bytes))
}
