package main

import (
	"encoding/json"
	"fmt"

	"github.com/icholy/utm"
	shp "github.com/julianhahn/shp_to_geojson"
)

func umt_to_latLng(point shp.GeoJSON_base_point) shp.GeoJSON_base_point {
	var zone utm.Zone = utm.Zone{Number: 52, Letter: 'S', North: true}
	if len(point) == 2 {
		var lat, lng float64 = zone.ToLatLon(point[0], point[1])
		return shp.GeoJSON_base_point{lng, lat}
	} else if len(point) == 3 {
		var lat, lng float64 = zone.ToLatLon(point[0], point[1])
		return shp.GeoJSON_base_point{lng, lat, point[2]}
	} else {
		return shp.GeoJSON_base_point{}
	}
}

func convert_to_latLng(object *shp.Feature) interface{} {
	json_string, _ := json.Marshal(object.Geometry)

	rightBounding := shp.GeoJSON_base_point{object.Properties["Xmax"].(float64), object.Properties["Ymax"].(float64), 0}
	bounding_latLng := umt_to_latLng(rightBounding)
	object.Properties["Xmax"] = bounding_latLng[0]
	object.Properties["Ymax"] = bounding_latLng[1]
	leftBounding := shp.GeoJSON_base_point{object.Properties["Xmin"].(float64), object.Properties["Ymin"].(float64), 0}
	bounding_latLng = umt_to_latLng(leftBounding)
	object.Properties["Xmin"] = bounding_latLng[0]
	object.Properties["Ymin"] = bounding_latLng[1]

	var geometry_type string = object.Geometry.(map[string]interface{})["type"].(string)
	switch geometry_type {
	case "MultiPoint":
		var multiPoint shp.GeoJSON_MultiPoint
		json.Unmarshal([]byte(json_string), &multiPoint)
		for i, point := range multiPoint.Coordinates {
			multiPoint.Coordinates[i] = umt_to_latLng(point)
		}
		object.Geometry = multiPoint
	case "MultiLineString":
		var multiLineString shp.GeoJSON_MultiLineString
		json.Unmarshal([]byte(json_string), &multiLineString)
		for i, line := range multiLineString.Coordinates {
			for j, point := range line {
				multiLineString.Coordinates[i][j] = umt_to_latLng(point)
			}
		}
		object.Geometry = multiLineString
	case "MultiPolygon":
		var multiPolygon shp.GeoJSON_MultiPolygon
		json.Unmarshal([]byte(json_string), &multiPolygon)
		for i, polygon := range multiPolygon.Coordinates {
			for j, line := range polygon {
				for k, point := range line {
					multiPolygon.Coordinates[i][j][k] = umt_to_latLng(point)
				}
			}
		}
		object.Geometry = multiPolygon
	case "Polygon":
		var polygon shp.GeoJSON_Polygon
		json.Unmarshal([]byte(json_string), &polygon)
		for i, line := range polygon.Coordinates {
			for j, point := range line {
				polygon.Coordinates[i][j] = umt_to_latLng(point)
			}
		}
		object.Geometry = polygon
	case "LineString":
		var lineString shp.GeoJSON_LineStrings
		json.Unmarshal([]byte(json_string), &lineString)
		for i, point := range lineString.Coordinates {
			lineString.Coordinates[i] = umt_to_latLng(point)
		}
		object.Geometry = lineString

	default:
		fmt.Print("on convert_to_latLng no type found")
	}
	return nil
}
