package main

import (
	"fmt"
	"os"
	"shp_to_geojson"
)

func main() {
	bytes, err := os.ReadFile("/Users/julian/react-babylon-starter/go/files/A1_NODE.shp")
	//bytes, err := os.ReadFile("/Users/julian/react-babylon-starter/go/files/B3_SURFACEMARK.shp")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(shp_to_geojson.ParseFromBytes(bytes))
}
