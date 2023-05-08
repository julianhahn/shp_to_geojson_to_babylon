import React, { useEffect, useRef } from "react";
import { useSelector } from "react-redux";

const CanvasElement = () => {

    let scene: BABYLON.Scene;
    let engine: BABYLON.Engine;
    const files = useSelector((state: any) => state.files);
    console.log(files);

    useEffect(() => {
        engine = new BABYLON.Engine(canvas.current, true);
        scene = new BABYLON.Scene(engine);
        scene.createDefaultCameraOrLight(true, true, true);
        let ground = BABYLON.MeshBuilder.CreateGround("ground", {
            width: 5,
            height: 5,
        }, scene);
        ground.position.y -= 1;
        let sphere = BABYLON.MeshBuilder.CreateSphere("Sphere", {
            diameter: 1,
        }, scene);
        if (scene.activeCamera) {
            let distance = 80;
            scene.activeCamera.position = new BABYLON.Vector3(distance, distance, distance);
        }
        engine.runRenderLoop(() => scene.render());
        return () => {
            scene.dispose();
            engine.dispose();
        }
    })

    let canvas = useRef(null);
    return (<canvas className={"babylon_canvas"} ref={canvas} />)
};

export default CanvasElement;