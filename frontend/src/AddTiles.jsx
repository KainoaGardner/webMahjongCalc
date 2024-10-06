import { useState } from "react";
import "./Hand.css";

function AddTiles() {
  const [tileCount, setTileCount] = useState(new Map());

  const [hand, setHand] = useState([]);
  function addTile(tile) {
    //add to tilecount
    if (hand.length < 14) {
      setTileCount(new Map(tileCount).set(tile, 1));
      setHand((h) => [...h, tile]);
    }
  }

  function removeTile(index) {
    setHand(hand.filter((_, i) => i !== index));
  }

  const manzu = ["M1", "M2", "M3", "M4", "M5", "M6", "M7", "M8", "M9", "M5A"];
  const pinzu = ["P1", "P2", "P3", "P4", "P5", "P6", "P7", "P8", "P9", "P5A"];
  const souzu = ["S1", "S2", "S3", "S4", "S5", "S6", "S7", "S8", "S9", "S5A"];
  const jihai = ["H1", "H2", "H3", "H4", "H5", "H6", "H7"];

  return (
    <>
      <div className="all">
        <div>
          {hand.map((tile, index) => (
            <img
              className="tile"
              key={index}
              src={`tiles/${tile}.png`}
              draggable="false"
              onClick={() => removeTile(index)}
            />
          ))}
        </div>

        <div>
          {manzu.map((tile) => (
            <img
              className="tile"
              key={tile}
              src={`tiles/${tile}.png`}
              draggable="false"
              onClick={() => addTile(tile)}
            />
          ))}
        </div>
        <div>
          {pinzu.map((tile) => (
            <img
              className="tile"
              key={tile}
              src={`tiles/${tile}.png`}
              draggable="false"
              onClick={() => addTile(tile)}
            />
          ))}
        </div>
        <div>
          {souzu.map((tile) => (
            <img
              className="tile"
              key={tile}
              src={`tiles/${tile}.png`}
              draggable="false"
              onClick={() => addTile(tile)}
            />
          ))}
        </div>
        <div>
          {jihai.map((tile) => (
            <img
              className="tile"
              key={tile}
              src={`tiles/${tile}.png`}
              draggable="false"
              onClick={() => addTile(tile)}
            />
          ))}
        </div>
      </div>
    </>
  );
}

export default AddTiles;
