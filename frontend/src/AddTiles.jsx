// import { useState } from "react";

import "./Hand.css";

function AddTiles({
  hand,
  setHand,
  chi,
  setChi,
  pon,
  setPon,
  kan,
  setKan,
  ankan,
  setAnkan,
  callType,
  setCallType,
  akaCall,
  tileCount,
  setTileCount,
  dora,
  setDora,
  uradora,
  setUradora,
  doraIndex,
  setDoraIndex,
  uradoraIndex,
  setUradoraIndex,
}) {
  function addTile(tile) {
    //add to tilecount
    const totalTiles =
      hand.length +
      chi.length +
      pon.length +
      (kan.length - parseInt(kan.length / 4)) +
      (ankan.length - parseInt(ankan.length / 4));

    if (totalTiles < 14) {
      const tiles = tileCount.get(tile);
      let akaTiles = tileCount.get(tile.concat("A"));
      if (!akaTiles) {
        akaTiles = 0;
      }

      if (
        tiles + akaTiles >= 4 ||
        (tile.length === 3 && tiles >= 1) ||
        (tile.length === 3 && tileCount.get(tile.slice(0, 2)) >= 4)
      ) {
        setDoraIndex(-1);
        setUradoraIndex(-1);
        return;
      }

      let tileAmount = 1;
      if (tiles) {
        tileAmount = tiles + 1;
      }

      if (doraIndex !== -1) {
        dora[doraIndex] = tile;
        setTileCount(new Map(tileCount).set(tile, tileAmount));
        setDora((d) => [...d]);
        setDoraIndex(-1);
      } else if (uradoraIndex !== -1) {
        setTileCount(new Map(tileCount).set(tile, tileAmount));
        uradora[uradoraIndex] = tile;
        setUradora((u) => [...u]);
        setUradoraIndex(-1);
      } else if (callType == "chi") {
        addChi(tile, totalTiles);
      } else if (callType == "pon") {
        addPon(tile, totalTiles);
      } else if (callType == "kan") {
        addKan(tile, totalTiles);
      } else if (callType == "ankan") {
        addAnkan(tile, totalTiles);
      } else {
        setTileCount(new Map(tileCount).set(tile, tileAmount));
        setHand((h) => [...h, tile]);
      }
    }
  }

  function addChi(tile, totalTiles) {
    if (
      totalTiles > 11 ||
      tile[1] > 7 ||
      tile[0] === "H" ||
      tile.length === 3
    ) {
      return;
    }

    if (akaCall && tile[1] !== "3" && tile[1] !== "4" && tile[1] !== "5") {
      return;
    }

    const tileNumber = tile[1] - "0";
    for (let i = 0; i < 3; i++) {
      const addTile = tile[0] + (tileNumber + i);
      let tileAmount = tileCount.get(addTile);
      let akaTileAmount = tileCount.get(addTile.concat("A"));
      if (!tileAmount) {
        tileAmount = 0;
      }
      if (!akaTileAmount) {
        akaTileAmount = 0;
      }
      if (
        tileAmount + akaTileAmount >= 4 ||
        (akaCall && tileNumber + i === 5 && akaTileAmount === 1)
      ) {
        setCallType("");
        return;
      }
    }

    for (let i = 0; i < 3; i++) {
      const addTile = tile[0] + (tileNumber + i);
      let tileAmount = tileCount.get(addTile);
      let akaTileAmount = tileCount.get(addTile.concat("A"));
      if (!tileAmount) {
        tileAmount = 0;
      }
      if (!akaTileAmount) {
        akaTileAmount = 0;
      }

      if (akaCall && tileNumber + i === 5) {
        chi.push(addTile.concat("A"));
        tileCount.set(addTile.concat("A"), akaTileAmount + 1);
      } else {
        chi.push(addTile);
        tileCount.set(addTile, tileAmount + 1);
      }
    }
    setChi((c) => [...c]);
    setTileCount(new Map(tileCount));
  }

  function addPon(tile, totalTiles) {
    if (totalTiles > 11 || (akaCall && tile[1] !== "5") || tile.length === 3) {
      return;
    }

    let tiles = 0;

    let akaTiles = tileCount.get(tile.concat("A"));
    if (tileCount.get(tile)) {
      tiles += tileCount.get(tile);
    }
    if (!akaTiles) {
      akaTiles = 0;
    }

    if (tiles + akaTiles > 1) {
      return;
    }
    pon.push(tile);
    pon.push(tile);
    if (akaCall) {
      pon.push(tile.concat("A"));
      tiles += 2;
      tileCount.set(tile.concat("A"), akaTiles + 1);
    } else {
      pon.push(tile);
      tiles += 3;
    }

    tileCount.set(tile, tiles);
    setPon((p) => [...p]);
    setTileCount(new Map(tileCount));
  }

  function addKan(tile, totalTiles) {
    if (totalTiles > 11 || (akaCall && tile[1] !== "5") || tile.length === 3) {
      return;
    }

    let tiles = tileCount.get(tile);
    if (!tiles) {
      tiles = 0;
    }

    let akaTiles = tileCount.get(tile.concat("A"));
    if (!akaTiles) {
      akaTiles = 0;
    }

    if (tiles + akaTiles > 0) {
      return;
    }

    kan.push(tile);
    kan.push(tile);
    kan.push(tile);
    if (akaCall) {
      kan.push(tile.concat("A"));
      tiles += 3;
      tileCount.set(tile.concat("A"), akaTiles + 1);
    } else {
      kan.push(tile);
      tiles += 4;
    }

    tileCount.set(tile, tiles);
    setKan((p) => [...p]);
    setTileCount(new Map(tileCount));
  }

  function addAnkan(tile, totalTiles) {
    if (totalTiles > 11 || (akaCall && tile[1] !== "5") || tile.length === 3) {
      return;
    }

    let tiles = tileCount.get(tile);
    if (!tiles) {
      tiles = 0;
    }

    let akaTiles = tileCount.get(tile.concat("A"));
    if (!akaTiles) {
      akaTiles = 0;
    }

    if (tiles + akaTiles > 0) {
      return;
    }

    ankan.push("B0");
    ankan.push(tile);
    if (akaCall) {
      ankan.push(tile.concat("A"));
      tiles += 3;
      tileCount.set(tile.concat("A"), akaTiles + 1);
    } else {
      ankan.push(tile);
      tiles += 4;
    }
    ankan.push("B0");

    tileCount.set(tile, tiles);
    setAnkan((k) => [...k]);
    setTileCount(new Map(tileCount));
  }

  function checkValidTile(tile) {
    if (
      tileCount.get(tile) >= 4 ||
      (tile.length === 3 && tileCount.get(tile) >= 1)
    ) {
      return false;
    }
    if (callType !== "" && tile.length === 3) {
      return false;
    }

    if (callType === "chi" && (tile[1] - "0" > 7 || tile[0] == "H")) {
      return false;
    }

    if (callType === "pon" && tileCount.get(tile) > 1) {
      return false;
    }

    if (
      akaCall &&
      callType === "chi" &&
      tile[1] !== "3" &&
      tile[1] !== "4" &&
      tile[1] !== "5"
    ) {
      return false;
    }

    if (
      (akaCall &&
        (callType === "pon" || callType === "kan" || callType === "ankan") &&
        tile[1] !== "5") ||
      (akaCall && tile[0] === "H" && callType !== "")
    ) {
      return false;
    }

    if (
      (callType === "kan" || callType === "ankan") &&
      tileCount.get(tile) > 0
    ) {
      return false;
    }

    return true;
  }

  const manzu = ["M1", "M2", "M3", "M4", "M5", "M6", "M7", "M8", "M9", "M5A"];
  const pinzu = ["P1", "P2", "P3", "P4", "P5", "P6", "P7", "P8", "P9", "P5A"];
  const souzu = ["S1", "S2", "S3", "S4", "S5", "S6", "S7", "S8", "S9", "S5A"];
  const jihai = ["H1", "H2", "H3", "H4", "H5", "H6", "H7"];

  return (
    <>
      <div>
        {manzu.map((tile) => (
          <img
            className={checkValidTile(tile) ? "tile" : "tile empty"}
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
            className={checkValidTile(tile) ? "tile" : "tile empty"}
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
            className={checkValidTile(tile) ? "tile" : "tile empty"}
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
            className={checkValidTile(tile) ? "tile" : "tile empty"}
            key={tile}
            src={`tiles/${tile}.png`}
            draggable="false"
            onClick={() => addTile(tile)}
          />
        ))}
      </div>
    </>
  );
}

export default AddTiles;
