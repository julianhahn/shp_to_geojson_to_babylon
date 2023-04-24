import { useEffect, useRef } from "react";
import * as BABYLON from "babylonjs";
import earcut from "earcut";

import "./App.css";

export const drawPolygon = (name: string, shape: BABYLON.Vector3[], scene: BABYLON.Scene) => {
  const polygon = BABYLON.MeshBuilder.CreatePolygon(
    name,
    {
      shape,
      sideOrientation: BABYLON.Mesh.DOUBLESIDE,
    },
    scene,
    earcut,
  );
  polygon.bakeCurrentTransformIntoVertices();
  return polygon;
};

function App() {
  let canvas = useRef(null);
  let engine: BABYLON.Engine;
  let scene: BABYLON.Scene;

  let background_colors = [
    "#B2BEB5",
    "#C5DCA0",
    "#F2E8C4",
    "#D4B996",
    "#C9A497",
    "#BFA5A4",
    "#D0B8A4",
    "#C7AC92",
    "#B4A7A4",
    "#D9C7B1",
    "#D4C7C6",
    "#C6B7A8",
    "#F6D0A6",
    "#D3AEB2",
    "#C9B2BB",
    "#D2B1C2",
    "#D0C6C7",
    "#B9C6CA",
    "#B5D9E5",
    "#C5B7C7",
  ];
  let line_colors = [
    "#F15A24",
    "#D40F0F",
    "#4D4D4D",
    "#2E2E2E",
    "#F9A03F",
    "#7CB490",
    "#4ECDC4",
    "#FF6B6B",
    "#1BBC9B",
    "#F7DC6F",
    "#9B59B6",
    "#34495E",
    "#E67E22",
    "#2980B9",
    "#2C3E50",
    "#E74C3C",
    "#16A085",
    "#8E44AD",
    "#27AE60",
    "#F1C40F",
  ];
  let point_colors = [
    "#FFA07A",
    "#98FB98",
    "#ADD8E6",
    "#FFC0CB",
    "#DDA0DD",
    "#FFB6C1",
    "#7FFFD4",
    "#87CEEB",
    "#F08080",
    "#FFDAB9",
    "#FF7F50",
    "#E6E6FA",
    "#B0E0E6",
    "#FFA500",
    "#20B2AA",
    "#AFEEEE",
    "#D8BFD8",
    "#FF6347",
    "#66CDAA",
    "#FFE4B5",
  ];

  useEffect(() => {
    engine = new BABYLON.Engine(canvas.current, true);
    scene = new BABYLON.Scene(engine);
    scene.createDefaultCameraOrLight(true, true, true);
    //scene.debugLayer.show();

    // Load the GeoJSON file using fetch or any other method you prefer
    fetch("test.geojson")
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

    let ground = BABYLON.MeshBuilder.CreateGround("ground", {
      width: 5,
      height: 5,
    });
    ground.position.y -= 1;
    let sphere = BABYLON.MeshBuilder.CreateSphere("Sphere", {
      diameter: 1,
    });
    if (scene.activeCamera) {
      let distance = 80;
      scene.activeCamera.position = new BABYLON.Vector3(distance, distance, distance);
    }
    engine.runRenderLoop(() => scene.render());
  });

  return (
    <div className="app">
      <canvas className="babylon_canvas" ref={canvas} id="babylonCanvas" />
    </div>
  );
}

export default App;
