 /*
    // Load the GeoJSON file using fetch or any other method you prefer
    fetch("localhost:8080/shapefiles")
      .then((response) => response.json())
      .then((data) => {
        let background_index = 0;
        let line_index = 0;
        let point_index = 0;
 
        // Loop through each feature in the GeoJSON file
        data.forEach((object: any) => {
          let colorIndex = Math.floor(Math.random() * background_colors.length);
          let color: any;
          var material = new BABYLON.StandardMaterial("material", scene);
          material.disableLighting = true;
 
          object.forEach((geometry: any) => {
            let points: any = [];
            if (geometry.type == "Polygon") {
              geometry.points.forEach((coord: any) => {
                points.push(new BABYLON.Vector3(coord[0], 0.1, coord[1]));
              });
              const polygon = drawPolygon("polygon", points, scene);
              color =
                background_colors[
                background_index < background_colors.length
                  ? background_index
                  : background_colors.length - 1
                ];
              polygon.material = material;
            }
            if (geometry.type == "LineString") {
              geometry.points.forEach((coord: any) => {
                points.push(new BABYLON.Vector3(coord[0], 0.2, coord[1]));
              });
              const line = BABYLON.MeshBuilder.CreateLines("lines", { points }, scene);
              color =
                line_colors[line_index < line_colors.length ? line_index : line_colors.length - 1];
              line.material = material;
            }
            if (geometry.type == "MultiPoint") {
              geometry.points.forEach((coord: any) => {
                points.push(new BABYLON.Vector3(coord[0], 0.3, coord[1]));
              });
              const disc = BABYLON.MeshBuilder.CreateDisc(
                "disc",
                { radius: 0.5, tessellation: 16 },
                scene,
              );
              disc.position = points[0];
              disc.rotation.x = Math.PI / 2;
              color =
                point_colors[
                point_index < point_colors.length ? point_index : point_colors.length - 1
                ];
              disc.material = material;
            }
 
            let babylon_color = BABYLON.Color3.FromHexString(color);
            material.diffuseColor = babylon_color;
            material.emissiveColor = babylon_color;
          });
        });
      });
 */