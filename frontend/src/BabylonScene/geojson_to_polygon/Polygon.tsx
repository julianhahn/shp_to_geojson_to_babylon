import React, { useEffect, useRef } from "react";
import { useSelector } from "react-redux";
import { Feature, FeatureCollection } from "geojson";
import * as BABYLON from "babylonjs";

type Props = {
  drawPolygon: (name: string, shape: BABYLON.Vector3[]) => BABYLON.Mesh;
  //drawPoint: (name: string, shape: BABYLON.Vector3[]) => BABYLON.Mesh;
  //drawLine: (name: string, shape: BABYLON.Vector3[]) => BABYLON.Mesh;
};

const Polygon = (props: Props) => {
  const collection: FeatureCollection = useSelector((state: any) => state.files.collection);

  useEffect(() => {
    console.log("features", JSON.stringify(collection));
    for (let feature of collection.features) {
      if (feature.geometry.type === "MultiPolygon") {
        for (let item of feature.geometry.coordinates) {
          for (let ring of item) {
            let vector_3 = ring.map((coordinate: any) => {
              return new BABYLON.Vector3(coordinate[0], coordinate[2], coordinate[1]);
            });
            let polygon = props.drawPolygon(feature.properties?.name, vector_3);
            console.log(polygon);
          }
        }
        // for (let i = 0; i < feature.geometry.coordinates.length; i++) {
        //   let vector_3 = feature.geometry.coordinates[i].map((coordinate: any) => {
        //     return new BABYLON.Vector3(coordinate[0], coordinate[2], coordinate[1]);
        //   });
        //   let polygon = props.drawPolygon(feature.properties.name, vector_3);
        //   polygon.position.y = 1;
        // }
      }
    }
  }, [collection]);

  return <></>;
};

export default Polygon;
