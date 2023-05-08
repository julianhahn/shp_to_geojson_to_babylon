import { configureStore, createSlice } from "@reduxjs/toolkit";
import { Feature } from "geojson";

const initialFeatures: { features: Feature[] } = {
  features: [],
};

export const geoJsonFeatures = createSlice({
  name: "files",
  initialState: initialFeatures,
  reducers: {
    setFeatures: (state, action) => {
      state.features = [...action.payload];
      return state;
    },
  },
});

export const { setFeatures } = geoJsonFeatures.actions;
export default geoJsonFeatures.reducer;

export const store = configureStore({
  reducer: {
    files: geoJsonFeatures.reducer,
  },
});
