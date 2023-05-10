import React, { useEffect, useState } from "react";
import "./index.css";
import { useDispatch } from "react-redux";
import { setCollection } from "../store";

const FilePicker = () => {
  const [files, setLocalFiles] = useState<File[] | null>(null);
  const dispatch = useDispatch();

  useEffect(() => {
    if (files && files.length > 0) {
      var formData = new FormData();
      for (let i = 0; i < files.length; i++) {
        formData.append("files", files[i]);
      }
      formData.append("WGS84", "false")
      fetch("http://localhost:8080/shapefiles", {
        method: "POST",
        body: formData,
      })
        .then((response) => {
          return response.json();
        })
        .then((data) => {
          dispatch(setCollection(data));
        })
        .catch((error) => {
          console.log(error);
        });
    }
  }, [files]);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file_list = e.target.files;
    if (file_list && file_list.length > 0) {
      var selected_files = [];
      for (let i = 0; i < file_list.length; i++) {
        selected_files.push(file_list[i]);
      }
      setLocalFiles(selected_files);
    }
  };

  return (
    <div className="filePickerContainer">
      <input type="file" accept=".shp" multiple onChange={handleFileChange} />
      {files?.map((file) => (
        <div key={file.name}>{file.name}</div>
      ))}
    </div>
  );
};

export { FilePicker };
