The goal is to display .shp (shapefile) data in BabylonJS. Since there is no direct adapter in BabylonJS we need to create the polygons ourself. Therefore we need the data from the shp file in a readable state.

The first step is to transform shp files to geojson.
I tried to do this in the browser and found more or less only those two packages:

- https://github.com/calvinmetcalf/shapefile-js
- https://www.npmjs.com/package/shp

the first one would not work and the second one only runs in node.
Since it seems easier in python, because the first hit on shp instantly worked we decided to put the coversion into the backend and handle the creation of the objects again in the frontend.

**Currently there is no rest api for the 'backend' and you have to run the script and copy a static file to the frontend but I'll update that if i get to it.**
