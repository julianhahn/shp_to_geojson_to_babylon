import React, { useEffect, useMemo, useRef, useState } from "react";
import * as BABYLON from "babylonjs";
import "./index.css";
import Polygon from "./geojson_to_polygon/Polygon";

const BabylonScene = () => {

  const [scene, setScene] = useState<BABYLON.Scene>();

  let engine: BABYLON.Engine;
  useEffect(() => {
    engine = new BABYLON.Engine(canvas.current, true);
    let scene = new BABYLON.Scene(engine);
    setScene(scene);
    scene.clearColor = BABYLON.Color4.FromHexString("#333333");

    scene.createDefaultCameraOrLight(true, true, true);
    if (scene.activeCamera) {
      let distance = 80;
      scene.activeCamera.position = new BABYLON.Vector3(distance, distance, distance);
    }

    engine.runRenderLoop(() => scene.render());
    return () => {
      scene.dispose();
      engine.dispose();
    };
  }, []);



  let canvas = useRef(null);
  return (
    <>
      <canvas className={"babylon_canvas"} ref={canvas} />
      <Polygon scn={scene} />
    </>
  );
};

export default BabylonScene;
