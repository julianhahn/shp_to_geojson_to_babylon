import os
import glob
import json
import shapefile
import argparse

# create a list of all the shapefiles in the ./files directory
all_geojson = []
edited_objects = []

# parse comand-line arguments
parser = argparse.ArgumentParser(description='Convert shapefiles to flattend json for frontend')
group = parser.add_mutually_exclusive_group(required=True)
group.add_argument("-f", "--file", help="Path to input shapefile")
group.add_argument("-d", "--dir", help="Path to input shapefile directory")
parser.add_argument("-o", "--output", help="Path to output JSON file")

args = parser.parse_args()

if args.dir:
    shapefiles = glob.glob(os.path.join(args.dir, "*.shp"))
elif args.file:
    shapefiles = [args.file]

for shapefile_path in shapefiles:
    sf = shapefile.Reader(shapefile_path)
    # get the geojson representation of the shapefile
    if (len(sf.shapes()) > 0):
        geojson = sf.shapes().__geo_interface__
        all_geojson.append(geojson)

min_value_y = None
min_value_x = None

# for object in all_geoJson
# coordinates = object["geometries"][0]
# for coordinate in coordinates
# and for point in coordinate


# loop over all objects / shapefiles
for object in all_geojson:
    for geometry in object["geometries"]:
        type = geometry["type"]
        coordinates = geometry["coordinates"]
        # if Polygon = A list of rings (each a list of (x, y) tuples)
        if (type == "Polygon" or type == "MultiLineString"):
            for coordinate in coordinates:
                for index, tuple in enumerate(coordinate):
                    if min_value_x is None or tuple[0] < min_value_x:
                        min_value_x = tuple[0]
                    if min_value_y is None or tuple[1] < min_value_y:
                        min_value_y = tuple[1]
        # if Point = A single (x, y) tuple
        elif (type == "Point"):
            if min_value_x is None or coordinates[0] < min_value_x:
                min_value_x = coordinates[0]
            if min_value_y is None or coordinates[1] < min_value_y:
                min_value_y = coordinates[1]
        # if MuliPoint A list of points (each a single (x, y) tuple)
        elif (type == "MultiPoint" or type == "LineString"):
            for point in coordinates:
                if min_value_x is None or point[0] < min_value_x:
                    min_value_x = point[0]
                if min_value_y is None or point[1] < min_value_y:
                    min_value_y = point[1]

# create a list of all objects. Every object has directly a list of geometries available with their type.
# no differentiation on which level the point nodes are lying

for object in all_geojson:
    flattend_geometries = []
    for geometry in object["geometries"]:
        new_coordinates_list = []
        type = geometry["type"]
        coordinates = geometry["coordinates"]
        # if Polygon = A list of rings (each a list of (x, y) tuples)
        if (type == "Polygon" or type == "MultiLineString"):
            for coordinate in coordinates:
                for index, tuple in enumerate(coordinate):
                    tuple = (tuple[0] - min_value_x,
                             tuple[1] - min_value_y)
                    new_coordinates_list.append(tuple)
        # if Point = A single (x, y) tuple
        elif (type == "Point"):
            new_point = (
                coordinates[0] - min_value_x, coordinates[1] - min_value_y)
            new_coordinates_list.append(new_point)
        # if MuliPoint A list of points (each a single (x, y) tuple)
        elif (type == "MultiPoint" or type == "LineString"):
            for index, point in enumerate(coordinates):
                new_point = (point[0]-min_value_x, point[1]-min_value_y)
                new_coordinates_list.append(new_point)

        flattend_geometries.append(
            {"type": type, "points": new_coordinates_list})

    edited_objects.append(flattend_geometries)


""" check if the user has provided an ouput path otherwise print back to the terminal """
if args.output:
    with open(args.output, "w") as f:
        json.dump(edited_objects, f)
else:
    print(edited_objects)
