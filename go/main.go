package main

import (
	"encoding/json"
	"fmt"
	"os"
	shp "shp_to_geojson"
)

func main() {
	var featureCollection shp.FeatureCollection = shp.FeatureCollection{
		Type:     "FeatureCollection",
		Features: []shp.Feature{},
	}
	// read all files in directory
	files, _ := os.ReadDir("/Users/julian/react-babylon-starter/go/files")

	for _, file := range files {
		if file.Name()[len(file.Name())-4:] == ".shp" {
			path := "/Users/julian/react-babylon-starter/go/files/" + file.Name()
			bytes, err := os.ReadFile(path)
			if err != nil {
				fmt.Println(err)
			}
			geoJSON, parseErr := shp.ParseFromBytes(bytes)
			if parseErr != nil {
				fmt.Println(parseErr)
			} else {
				var newFeature shp.Feature = shp.Feature{
					Type:       "Feature",
					Properties: map[string]string{},
					Geometry:   convert_to_latLng(geoJSON),
				}
				featureCollection.Features = append(featureCollection.Features, newFeature)
			}
		}
	}
	var content, err = json.Marshal(featureCollection)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}
