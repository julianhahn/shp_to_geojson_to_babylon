import earcut from "earcut";
import * as BABYLON from "babylonjs";
import { point_colors, background_colors, line_colors } from "../../data";

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
  //polygon.scaling = new BABYLON.Vector3(100000, 100000, 100000);
  polygon.bakeCurrentTransformIntoVertices();
  return polygon;
};

const drawLine = (name: string, shape: BABYLON.Vector3[], scene: BABYLON.Scene) => {
  const line = BABYLON.MeshBuilder.CreateLines(name, { points: shape }, scene);
  //line.scaling = new BABYLON.Vector3(100000, 100000, 100000);
  line.enableEdgesRendering();
  line.edgesWidth = 20.0;
  line.bakeCurrentTransformIntoVertices();
  return line;
};

const drawPoint = (name: string, position: BABYLON.Vector3, scene: BABYLON.Scene) => {
  const point = BABYLON.MeshBuilder.CreateDisc(name, { radius: 0.5, tessellation: 16 }, scene);
  point.position = position;
  point.rotation.x = Math.PI / 2;
  //point.scaling = new BABYLON.Vector3(100000, 100000, 100000);
  point.bakeCurrentTransformIntoVertices();
  return point;
};

export { drawPolygon, drawLine, drawPoint };
