import { useState } from "react";

import "./Hand.css";

function AddTiles() {
  const [tileCount, setTileCount] = useState(new Map());

  const [hand, setHand] = useState([]);
  function addTile(tile) {
    //add to tilecount
    if (hand.length < 14) {
      const tiles = tileCount.get(tile);
      let akaTiles = tileCount.get(tile.concat("A"));
      if (!akaTiles) {
        akaTiles = 0;
      }

      if (tiles + akaTiles >= 4 || (tile.length === 3 && tiles >= 1)) {
        return;
      }

      if (tile.length === 3 && tileCount.get(tile.slice(0, 2)) >= 4) {
        return;
      }

      let tileAmount = 1;
      if (tiles) {
        tileAmount = tiles + 1;
      }

      setTileCount(new Map(tileCount).set(tile, tileAmount));
      setHand((h) => [...h, tile]);
    }
  }

  function removeTile(tile, index) {
    setTileCount(new Map(tileCount).set(tile, tileCount.get(tile) - 1));
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
              onClick={() => removeTile(tile, index)}
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
