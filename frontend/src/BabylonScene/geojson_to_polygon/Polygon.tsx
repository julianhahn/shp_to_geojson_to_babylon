import { useEffect, useState, } from "react";
import { useSelector } from "react-redux";
import { FeatureCollection } from "geojson";
import { drawPolygon, drawLine, drawPoint } from "./mesh_functions";
import * as BABYLON from "babylonjs";
import { line_colors, background_colors, point_colors } from "../../data";

type Props = {
  scn: BABYLON.Scene | undefined;
};



const Polygon = (props: Props) => {
  const collection: FeatureCollection = useSelector((state: any) => state.files.collection);

  const [background_index, setBackground] = useState(7);
  const [line_index, setLine] = useState(5);
  const [point_index, setPoint] = useState(10);

  const [meshes, setMeshes] = useState<BABYLON.Mesh[]>([]);

  useEffect(() => {
    if (meshes.length > 0) {
      const groupNode = new BABYLON.TransformNode("groupNode", props.scn);
      for (let mesh of meshes) {
        mesh.parent = groupNode;
      }
      const center = BABYLON.Vector3.Zero();
      meshes.forEach(mesh => {
        center.addInPlace(mesh.getBoundingInfo().boundingBox.center);
      });
      center.scaleInPlace(1 / meshes.length);
      groupNode.position = center.scale(-1);
    }
  }, [meshes])


  var temp_meshes: BABYLON.Mesh[] = [];
  useEffect(() => {
    if (!props.scn) return;
    const scene = props.scn;

    if (meshes.length > 0) {
      for (let mesh of meshes) {
        mesh.dispose();
      }
      setMeshes([]);
    }

    var ground_material = scene.getMaterialByName("ground") as BABYLON.StandardMaterial;
    if (scene.getMaterialByName("ground") === null) {
      ground_material = new BABYLON.StandardMaterial("ground", scene);
      ground_material.disableLighting = true;
      ground_material.diffuseColor = BABYLON.Color3.FromHexString(background_colors[background_index]);
      ground_material.emissiveColor = BABYLON.Color3.FromHexString(background_colors[background_index]);
    }

    var line_material = scene.getMaterialByName("line") as BABYLON.StandardMaterial;
    if (scene.getMaterialByName("line") === null) {
      line_material = new BABYLON.StandardMaterial("line", scene);
      line_material.disableLighting = true;
      line_material.diffuseColor = BABYLON.Color3.FromHexString(line_colors[line_index]);
      line_material.emissiveColor = BABYLON.Color3.FromHexString(line_colors[line_index]);
    }

    var point_material = scene.getMaterialByName("point") as BABYLON.StandardMaterial;
    if (scene.getMaterialByName("point") === null) {
      point_material = new BABYLON.StandardMaterial("point", scene);
      point_material.disableLighting = true;
      point_material.diffuseColor = BABYLON.Color3.FromHexString(point_colors[point_index]);
      point_material.emissiveColor = BABYLON.Color3.FromHexString(point_colors[point_index]);
    }



    let [Xmax, Xmin, Ymax, Ymin, Zmax, Zmin] = [Number.MIN_VALUE, Number.MAX_VALUE, Number.MIN_VALUE, Number.MAX_VALUE, Number.MIN_VALUE, Number.MAX_VALUE]
    // find max value of all objects to subtract from all values
    for (let feature of collection.features) {
      if (feature.properties) {
        const [f_Xmax, f_Xmin, f_Ymax, f_Ymin, f_Zmax, f_Zmin] = [feature.properties.Xmax, feature.properties.Xmin, feature.properties.Ymax, feature.properties.Ymin, feature.properties.Zmax, feature.properties.Zmin]
        if (f_Xmax > Xmax) Xmax = f_Xmax
        if (f_Xmin < Xmin) Xmin = f_Xmin
        if (f_Ymax > Ymax) Ymax = f_Ymax
        if (f_Ymin < Ymin) Ymin = f_Ymin
        if (f_Zmax > Zmax) Zmax = f_Zmax
        if (f_Zmin < Zmin) Zmin = f_Zmin
      }
    }

    for (const feature of collection.features) {
      if (feature.geometry.type === "MultiPolygon") {
        for (let polygon of feature.geometry.coordinates) {
          for (let ring of polygon) {
            let vector_3 = ring.map((coordinate: any) => {
              return new BABYLON.Vector3(coordinate[0] - Xmin, coordinate[2] - Zmin - 0.1, coordinate[1] - Ymin);
            });
            if (vector_3.length > 2) {
              let polygon_mesh = drawPolygon(feature.properties?.name ?? "", vector_3, scene);
              polygon_mesh.material = ground_material;
              temp_meshes.push(polygon_mesh);
            }
          }
        }
      } else if (feature.geometry.type === "MultiLineString") {
        for (let line of feature.geometry.coordinates) {
          let vector_3 = line.map((coordinate: any) => {
            return new BABYLON.Vector3(coordinate[0] - Xmin, coordinate[2] - Zmin + 0.1, coordinate[1] - Ymin);
          });
          let line_mesh = drawLine(feature.properties?.name, vector_3, scene);
          line_mesh.material = line_material;
          line_mesh.edgesColor = BABYLON.Color4.FromHexString(line_colors[line_index]);
          temp_meshes.push(line_mesh);
        }
      } else if (feature.geometry.type === "LineString") {
        let vector_3 = feature.geometry.coordinates.map((coordinate: any) => {
          return new BABYLON.Vector3(coordinate[0] - Xmin, coordinate[2] - Zmin + 0.1, coordinate[1] - Ymin);
        });
        let line_mesh = drawLine(feature.properties?.name, vector_3, scene);
        line_mesh.material = line_material;
        temp_meshes.push(line_mesh);
      }
      else if (feature.geometry.type === "Point") {
        feature.geometry.coordinates.map((coordinate: any) => {
          let point = new BABYLON.Vector3(coordinate[0] - Xmin, coordinate[2] - Zmin, coordinate[1] - Ymin);
          let point_mesh = drawPoint(feature.properties?.name, point, scene);
          point_mesh.material = point_material;
          temp_meshes.push(point_mesh);
        });
        //props.drawPoint(feature.properties?.name, vector_3);
      } else if (feature.geometry.type === "MultiPoint") {
        for (let point of feature.geometry.coordinates) {
          let vector_3 = new BABYLON.Vector3(point[0] - Xmin, point[2] - Zmin + 0.2, point[1] - Ymin)
          if (vector_3.y > 3) {
            const sphere_mesh = BABYLON.MeshBuilder.CreateSphere("sphere", { diameter: 0.6 }, scene);
            sphere_mesh.position = vector_3;
            sphere_mesh.material = point_material;
            temp_meshes.push(sphere_mesh);
          } else {
            let point_mesh = drawPoint(feature.properties?.name, vector_3, scene);
            point_mesh.material = point_material;
            temp_meshes.push(point_mesh);
          }
        }
      }
    }

    setMeshes(temp_meshes);
  }, [collection, background_index, line_index, point_index]);

  return (
    <input type="slider" max={background_colors.length} onChange={(e) => {
      setBackground(parseInt(e.target.value));
    }} />
  );
};

export default Polygon;
