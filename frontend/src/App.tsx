import { store } from "./store";
import { Provider } from "react-redux";
import "./App.css";
//import { line_colors, background_colors, point_colors } from "./data";
import { FilePicker } from "./FilePicker/FilePicker"
import BabylonScene from "./BabylonScene/bablyonScene";

function App() {

  return (
    <Provider store={store}>
      <div className="app">
        <FilePicker />
        <BabylonScene />
      </div>
    </Provider>
  );
}

export default App;
