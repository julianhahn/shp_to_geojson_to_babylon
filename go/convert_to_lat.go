package main

import (
	"encoding/json"
	"regexp"
	shp "shp_to_geojson"

	"github.com/icholy/utm"
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

func convert_to_latLng(geojson_object_string string) interface{} {
	r := regexp.MustCompile(`"type"\s*:\s*"(\w+)"`)
	match := r.FindStringSubmatch(geojson_object_string)
	if len(match) > 1 {
		switch match[1] {
		case "MultiPoint":
			var geojson_object shp.GeoJSON_MultiPoint
			json.Unmarshal([]byte(geojson_object_string), &geojson_object)
			for i, point := range geojson_object.Coordinates {
				geojson_object.Coordinates[i] = umt_to_latLng(point)
			}
			return geojson_object
		case "LineString":
			var geojson_object shp.GeoJSON_LineStrings
			json.Unmarshal([]byte(geojson_object_string), &geojson_object)
			for i, point := range geojson_object.Coordinates {
				geojson_object.Coordinates[i] = umt_to_latLng(point)
			}
			return geojson_object
		case "MultiLineString":
			var geojson_object shp.GeoJSON_MultiLineString
			json.Unmarshal([]byte(geojson_object_string), &geojson_object)
			for i, line := range geojson_object.Coordinates {
				for j, point := range line {
					geojson_object.Coordinates[i][j] = umt_to_latLng(point)
				}
			}
			return geojson_object
		case "Polygon":
			var geojson_object shp.GeoJSON_Polygon
			json.Unmarshal([]byte(geojson_object_string), &geojson_object)
			for i, line := range geojson_object.Coordinates {
				for j, point := range line {
					geojson_object.Coordinates[i][j] = umt_to_latLng(point)
				}
			}
			return geojson_object
		case "MultiPolygon":
			var geojson_object shp.GeoJSON_MultiPolygon
			json.Unmarshal([]byte(geojson_object_string), &geojson_object)
			for i, polygon := range geojson_object.Coordinates {
				for j, line := range polygon {
					for k, point := range line {
						geojson_object.Coordinates[i][j][k] = umt_to_latLng(point)
					}
				}
			}
			return geojson_object
		default:
			return nil
		}
	}
	return nil
}
