import React, { useEffect, useMemo, useRef } from "react";
import * as BABYLON from "babylonjs";
import earcut from "earcut";
import "./index.css";
import CanvasElement from "./CanvasElement/CanvasElement";

const BabylonScene = () => {

  return (
    <CanvasElement />
  );
};

const drawPolygon = (name: string, shape: BABYLON.Vector3[], scene: BABYLON.Scene) => {
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

export default BabylonScene;
export { drawPolygon, BabylonScene };
