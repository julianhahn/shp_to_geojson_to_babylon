package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	shp "shp_to_geojson"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "*",
	}))

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

		var should_convert bool = form.Value["WGS84"][0] == "true"

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
			return c.SendString(content)

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
				newFeature := shp.Feature{}
				json.Unmarshal([]byte(content), &newFeature)
				if should_convert {
					convert_to_latLng(&newFeature)
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

	log.Fatal(app.Listen(":8080"))
}
