import "./App.css";
// import AddTiles from "./AddTiles.jsx";
import Hand from "./Hand.jsx";
import Information from "./Information.jsx";

function App() {
  return (
    <>
      <h1 className="appTitle">Riichi Mahjong Hand Calculator</h1>
      <hr />
      <Hand />
      <Information />
    </>
  );
}

export default App;
