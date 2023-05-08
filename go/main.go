package main

import (
	"fmt"
	"io/ioutil"

	shp "shp_to_geojson"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Post("/shapefiles", func(c *fiber.Ctx) error {

		form, formErr := c.MultipartForm()
		if form.File["files"] == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Missing files field",
			})
		} else if formErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Error parsing form: %v", formErr),
			})
		}
		files := form.File["files"]

		if len(files) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "postes files were empty",
			})
		} else if len(files) == 1 {
			file, openError := files[0].Open()
			if openError != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": fmt.Sprintf("Error opening file: %v", openError),
				})
			}
			shapefileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"file":  fmt.Sprintf("%v", files[0].Filename),
					"error": fmt.Sprintf("Error opening file: %v", err),
				})
			}
			content, parseErr := shp.ParseFromBytes(shapefileBytes)
			if parseErr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"file":  fmt.Sprintf("%v", files[0].Filename),
					"error": fmt.Sprintf("Error parsing file: %v", parseErr),
				})
			}
			return c.JSON(content)

		} else if len(files) > 1 {
			var featureCollection shp.FeatureCollection = shp.FeatureCollection{
				Type:     "FeatureCollection",
				Features: []shp.Feature{},
			}

			for _, file_ref := range files {
				file, openError := file_ref.Open()
				if openError != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": fmt.Sprintf("Error opening file: %v", openError),
					})
				}
				shapefileBytes, err := ioutil.ReadAll(file)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"file":  fmt.Sprintf("%v", file_ref.Filename),
						"error": fmt.Sprintf("Error opening file: %v", err),
					})
				}
				content, parseErr := shp.ParseFromBytes(shapefileBytes)
				if parseErr != nil {
					if parseErr.Error() == "no content found in shp file" {
						continue
					}
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"file":  fmt.Sprintf("%v", file_ref.Filename),
						"error": fmt.Sprintf("Error parsing file: %v", parseErr),
					})
				}
				var newFeature shp.Feature = shp.Feature{
					Type:       "Feature",
					Properties: map[string]string{},
					Geometry:   convert_to_latLng(content),
				}
				featureCollection.Features = append(featureCollection.Features, newFeature)
			}
			return c.JSON(featureCollection)
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Something went wrong",
			})
		}

	})
	app.Listen(":8080")

	/* if len(part.FileName()) > 4 && part.FileName()[len(part.FileName())-4:] == ".shp" {
		fileBytes, fileReadError := ioutil.ReadAll(part)
		if fileReadError != nil {
			fmt.Println(fileReadError)
		}

		geoJSON, parseErr := shp.ParseFromBytes(fileBytes)
		if parseErr != nil {
			fmt.Println(parseErr)
		}
		fmt.Fprintf(w, "data: %v", geoJSON)

	} */

	/*
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
		fmt.Println(string(content)) */
}
