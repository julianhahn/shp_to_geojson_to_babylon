package shp_to_json

import (
	"fmt"
	"os"

	"github.com/jonas-p/go-shp"
)

func Convert(inputFilePath string) string {

	shape, err := shp.Open("path/to/shapefile.shp")
	if err != nil {
		fmt.Println("Error opening shapefile:", err)
		os.Exit(1)
	}
	defer shape.Close()

	// read all the records in the shapefile
	for shape.Next() {
		n, shapeData := shape.Shape()

		// create a new GeoJSON feature based on the shape
		feature := geojson.NewFeature(nil)
		switch shapeType {
		case shp.POINT, shp.POINTM, shp.POINTZ:
			point := geojson.NewPointGeometry([]float64{shapeData.X, shapeData.Y})
			feature.SetGeometry(point)
		case shp.POLYLINE, shp.POLYLINEM, shp.POLYLINEZ:
			lineString := geojson.NewLineStringGeometry([][]float64{shapeData.Points})
			feature.SetGeometry(lineString)
		case shp.POLYGON, shp.POLYGONM, shp.POLYGONZ:
			polygon := geojson.NewPolygonGeometry([][][]float64{shapeData.Polygons})
			feature.SetGeometry(polygon)
		}

		// print out the GeoJSON for the feature
		geojsonString, err := feature.MarshalJSON()
		if err != nil {
			fmt.Println("Error creating GeoJSON for feature:", err)
			os.Exit(1)
		}
		fmt.Println(string(geojsonString))
	}
	return inputFilePath
}
