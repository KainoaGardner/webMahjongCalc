import "./App.css";
import Hand from "./Hand.jsx";
import Yaku from "./Yaku.jsx";
import Fu from "./Fu.jsx";
import About from "./About.jsx";
import { useState } from "react";

function App() {
  const [page, setPage] = useState("home"); //hand yaku fu about

  return (
    <>
      <div className="nav">
        <button
          className={page === "home" ? "on navButton" : "navButton"}
          onClick={() => setPage("home")}
        >
          Home
        </button>
        <button
          className={page === "yaku" ? "on navButton" : "navButton"}
          onClick={() => setPage("yaku")}
        >
          Yaku
        </button>
        <button
          className={page === "fu" ? "on navButton" : "navButton"}
          onClick={() => setPage("fu")}
        >
          Fu/Scoring
        </button>
        <button
          className={page === "about" ? "on navButton" : "navButton"}
          onClick={() => setPage("about")}
        >
          About
        </button>
      </div>
      <Tab page={page} />
    </>
  );
}

function Tab({ page }) {
  switch (page) {
    case "yaku":
      return <Yaku />;
    case "fu":
      return <Fu />;

    case "about":
      return <About />;

    default:
      return (
        <>
          <h1 className="appTitle">Riichi Mahjong Hand Calculator</h1>
          <hr />
          <Hand />
        </>
      );
  }
}

export default App;
