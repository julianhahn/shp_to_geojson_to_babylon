import os
import glob
import json
import shapefile

# create a list of all the shapefiles in the ./files directory
shapefiles = glob.glob("./files/*.shp")
sf = shapefile.Reader("./files/B3_SURFACEMARK")
# get the geojson representation of the shapefile
geojson = sf.shapes().__geo_interface__

# loop over all the geometries in the geojson
min_value_y = None
min_value_x = None
for geometry in geojson['geometries']:
    # loop over all the coordinates in the geometry
    for coordinate in geometry['coordinates']:
        # loop over all the tuples in the coordinate
        for point in coordinate:
            # check if the min_value is None or if the point is smaller than the current min_value
            if min_value_x is None or point[0] < min_value_x:
                min_value_x = point[0]
            if min_value_y is None or point[1] < min_value_y:
                min_value_y = point[1]


for geometry in geojson['geometries']:
    # loop over all the coordinates in the geometry
    for coordinate in geometry['coordinates']:
        # loop over all the tuples in the coordinate
        for i in range(len(coordinate)):
            # subtract the y value of the tuple with the min_value
            coordinate[i] = (coordinate[i][0] - min_value_x,
                             coordinate[i][1] - min_value_y)
            print(coordinate[i])
# write the updated geojson to a file
with open("test.geojson", "w") as f:
    json.dump(geojson, f)
