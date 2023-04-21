import { useEffect, useRef } from "react";
import * as BABYLON from "babylonjs";
import "./App.css";

function App() {
  let canvas = useRef(null);
  let engine: BABYLON.Engine;
  let scene: BABYLON.Scene;

  useEffect(() => {
    engine = new BABYLON.Engine(canvas.current, true);
    scene = new BABYLON.Scene(engine);
    scene.createDefaultCameraOrLight(true, true, true);
    let ground = BABYLON.MeshBuilder.CreateGround("ground", {
      width: 5,
      height: 5,
    });
    ground.position.y -= 1;
    let sphere = BABYLON.MeshBuilder.CreateSphere("Sphere", {
      diameter: 1,
    });
    if (scene.activeCamera) {
      scene.activeCamera.position = new BABYLON.Vector3(6, 6, 6);
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
