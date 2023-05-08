import { configureStore, createSlice } from "@reduxjs/toolkit";

const initialFiles: { files: Blob[] } = {
  files: [],
};

export const filesSlice = createSlice({
  name: "files",
  initialState: initialFiles,
  reducers: {
    setFiles: (state, action) => {
      state.files = [...action.payload];
      return state;
    },
  },
});

export const { setFiles } = filesSlice.actions;
export default filesSlice.reducer;

export const store = configureStore({
  reducer: {
    files: filesSlice.reducer,
  },
});
