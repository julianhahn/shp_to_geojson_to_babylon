import React, { useEffect, useMemo, useRef } from "react";
import * as BABYLON from "babylonjs";
import earcut from "earcut";
import "./index.css";
import CanvasElement from "./geojson_to_polygon/Polygon";
import Polygon from "./geojson_to_polygon/Polygon";

const BabylonScene = () => {
  let scene: BABYLON.Scene;
  let engine: BABYLON.Engine;

  const drawPolygon = (name: string, shape: BABYLON.Vector3[]) => {
    console.log(shape);

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

  useEffect(() => {
    engine = new BABYLON.Engine(canvas.current, true);
    scene = new BABYLON.Scene(engine);
    scene.createDefaultCameraOrLight(true, true, true);
    let ground = BABYLON.MeshBuilder.CreateGround(
      "ground",
      {
        width: 5,
        height: 5,
      },
      scene,
    );
    ground.position.y -= 1;
    let sphere = BABYLON.MeshBuilder.CreateSphere(
      "Sphere",
      {
        diameter: 1,
      },
      scene,
    );
    if (scene.activeCamera) {
      let distance = 80;
      scene.activeCamera.position = new BABYLON.Vector3(distance, distance, distance);
    }

    engine.runRenderLoop(() => scene.render());
    return () => {
      scene.dispose();
      engine.dispose();
    };
  });

  let canvas = useRef(null);
  return (
    <>
      <canvas className={"babylon_canvas"} ref={canvas} />
      <Polygon drawPolygon={drawPolygon} />
    </>
  );
};

export default BabylonScene;
