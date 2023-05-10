import { configureStore, createSlice } from "@reduxjs/toolkit";
import { FeatureCollection } from "geojson";

const initialFeatures: { collection: FeatureCollection; objects: any } = {
  collection: {
    type: "FeatureCollection",
    features: [],
  },
  objects: {},
};

export const geoJsonFeatures = createSlice({
  name: "files",
  initialState: initialFeatures,
  reducers: {
    setFeatures: (state, action) => {
      state.collection.features = [...action.payload];
      return state;
    },
    setCollection: (state, action) => {
      state.collection = action.payload;
      return state;
    },
  },
});

export const { setFeatures, setCollection } = geoJsonFeatures.actions;
export default geoJsonFeatures.reducer;

export const store = configureStore({
  reducer: {
    files: geoJsonFeatures.reducer,
  },
});
