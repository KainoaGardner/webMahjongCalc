import { useState } from "react";

import AddTiles from "./AddTiles.jsx";
import Dora from "./Dora.jsx";
import ScoringOptions from "./ScoringOptions.jsx";
import Results from "./Results.jsx";
import Tenpai from "./Tenpai.jsx";

import "./Hand.css";

function Hand() {
  const [tileCount, setTileCount] = useState(new Map());

  const [hand, setHand] = useState([]);
  const [chi, setChi] = useState([]);
  const [pon, setPon] = useState([]);
  const [kan, setKan] = useState([]);
  const [ankan, setAnkan] = useState([]);
  const [callType, setCallType] = useState("");
  const [akaCall, setAkaCall] = useState();

  const [dora, setDora] = useState(["B0", "B0", "B0", "B0"]);
  const [uradora, setUradora] = useState(["B0", "B0", "B0", "B0"]);
  const [doraIndex, setDoraIndex] = useState(-1);
  const [uradoraIndex, setUradoraIndex] = useState(-1);

  const [oya, setOya] = useState(true); //true oya false ko
  const [agari, setAgari] = useState(true); //true ron false tsumo
  const [riichi, setRiichi] = useState("None"); //None riichi wriichi
  const [ippatsu, setIppatsu] = useState(false); //true ippatsu false none
  const [chankan, setChankan] = useState(false); //true chankan false none
  const [rinshan, setRinshan] = useState(false); //true rinshan false none
  const [haitei, setHaitei] = useState("None"); //None haitei houtei
  const [tenhou, setTenhou] = useState("None"); //None tenhou chiihou
  const [honba, setHonba] = useState(0); //0 >
  const [riichibou, setRiichibou] = useState(0); //0 >
  const [kiriage, setKiriage] = useState(false); //true kiriage false none
  const [bakaze, setBakaze] = useState("H1"); //H1 ton H2 nan H3 sha H4 pei
  const [jikaze, setJikaze] = useState("H1"); //H1 ton H2 nan H3 sha H4 pei

  function changeCallType(type) {
    if (callType === type) {
      setCallType("");
    } else {
      setCallType(type);
    }
  }

  function removeTile(tile, index) {
    setTileCount(new Map(tileCount).set(tile, tileCount.get(tile) - 1));
    setHand(hand.filter((_, i) => i !== index));
  }

  function removeChi(tile, index) {
    let startIndex = parseInt(index / 3) * 3;
    for (let i = 0; i < 3; i++) {
      tileCount.set(chi[i], tileCount.get(chi[i] - 1));
    }

    chi.splice(startIndex, 3);
    setTileCount(new Map(tileCount));
    setChi((c) => [...c]);
  }

  function removePon(tile, index) {
    let startIndex = parseInt(index / 3) * 3;

    // setTileCount(new Map(tileCount).set(tile, tileCount.get(tile) - 1));
    if (tile.length === 3) {
      tile = tile.slice(0, -1);
    }

    if (tile[1] !== "5") {
      const tileAmount = tileCount.get(tile);
      setTileCount(new Map(tileCount).set(tile, tileAmount - 3));
    } else {
      let akaFound = false;
      for (let i = 0; i < 3; i++) {
        if (pon[i].length === 3) {
          akaFound = true;
        }
      }
      if (akaFound) {
        const tileAmount = tileCount.get(tile);
        const akaTileAmount = tileCount.get(tile.concat("A"));
        tileCount.set(tile, tileAmount - 2);
        tileCount.set(tile.concat("A"), akaTileAmount - 1);
        setTileCount(new Map(tileCount));
      } else {
        const tileAmount = tileCount.get(tile);
        setTileCount(new Map(tileCount).set(tile, tileAmount - 3));
      }
    }

    pon.splice(startIndex, 3);
    setPon((p) => [...p]);
  }

  function removeKan(tile, index) {
    let startIndex = parseInt(index / 4) * 4;

    if (tile.length === 3) {
      tile = tile.slice(0, -1);
    }

    if (tile[1] !== "5") {
      const tileAmount = tileCount.get(tile);
      setTileCount(new Map(tileCount).set(tile, tileAmount - 4));
    } else {
      let akaFound = false;
      for (let i = 0; i < 4; i++) {
        if (kan[i].length === 4) {
          akaFound = true;
        }
      }
      if (akaFound) {
        const tileAmount = tileCount.get(tile);
        const akaTileAmount = tileCount.get(tile.concat("A"));
        tileCount.set(tile, tileAmount - 3);
        tileCount.set(tile.concat("A"), akaTileAmount - 1);
        setTileCount(new Map(tileCount));
      } else {
        const tileAmount = tileCount.get(tile);
        setTileCount(new Map(tileCount).set(tile, tileAmount - 4));
      }
    }

    kan.splice(startIndex, 4);
    setKan((k) => [...k]);
  }

  function removeAnkan(tile, index) {
    let startIndex = parseInt(index / 4) * 4;

    if (tile.length === 3 || tile === "B0") {
      tile = ankan[startIndex + 1];
    }

    if (tile[1] !== "5") {
      const tileAmount = tileCount.get(tile);
      setTileCount(new Map(tileCount).set(tile, tileAmount - 4));
    } else {
      let akaFound = false;
      for (let i = 0; i < 4; i++) {
        if (ankan[i].length === 4) {
          akaFound = true;
        }
      }
      if (akaFound) {
        const tileAmount = tileCount.get(tile);
        const akaTileAmount = tileCount.get(tile.concat("A"));
        tileCount.set(tile, tileAmount - 3);
        tileCount.set(tile.concat("A"), akaTileAmount - 1);
        setTileCount(new Map(tileCount));
      } else {
        const tileAmount = tileCount.get(tile);
        setTileCount(new Map(tileCount).set(tile, tileAmount - 4));
      }
    }

    ankan.splice(startIndex, 4);
    setAnkan((k) => [...k]);
  }

  function clearHand() {
    setCallType("");
    for (let i = hand.length; i >= 0; i--) {
      tileCount.set(hand[i], tileCount.get(hand[i]) - 1);
    }
    for (let i = chi.length; i >= 0; i--) {
      tileCount.set(chi[i], tileCount.get(chi[i]) - 1);
    }
    for (let i = pon.length; i >= 0; i--) {
      tileCount.set(pon[i], tileCount.get(pon[i]) - 1);
    }
    for (let i = kan.length; i >= 0; i--) {
      tileCount.set(kan[i], tileCount.get(kan[i]) - 1);
    }
    for (let i = ankan.length; i >= 0; i--) {
      if (i % 4 === 0) {
        tileCount.set(ankan[i + 1], tileCount.get(ankan[i + 1]) - 1);
      } else if (i % 4 === 3) {
        tileCount.set(ankan[i - 2], tileCount.get(ankan[i - 2]) - 1);
      }
      tileCount.set(ankan[i], tileCount.get(ankan[i]) - 1);
    }

    setTileCount(new Map(tileCount));

    setHand([]);
    setChi([]);
    setPon([]);
    setKan([]);
    setAnkan([]);
  }

  function clearDora() {
    for (let i = dora.length; i >= 0; i--) {
      if (dora[i] !== "B0") {
        tileCount.set(dora[i], tileCount.get(dora[i]) - 1);
      }
    }
    for (let i = uradora.length; i >= 0; i--) {
      if (uradora[i] !== "B0") {
        tileCount.set(uradora[i], tileCount.get(uradora[i]) - 1);
      }
    }

    setTileCount(new Map(tileCount));

    setDora(["B0", "B0", "B0", "B0"]);
    setUradora(["B0", "B0", "B0", "B0"]);
  }

  return (
    <>
      <div className="mainBody">
        <div className="leftMain">
          <div>
            {hand.length +
              chi.length +
              pon.length +
              (kan.length - parseInt(kan.length / 4)) +
              (ankan.length - parseInt(ankan.length / 4)) ===
              14 && (
              <Results
                hand={hand}
                chi={chi}
                pon={pon}
                kan={kan}
                ankan={ankan}
                oya={oya}
                agari={agari}
                riichi={riichi}
                ippatsu={ippatsu}
                chankan={chankan}
                rinshan={rinshan}
                haitei={haitei}
                tenhou={tenhou}
                honba={honba}
                riichibou={riichibou}
                kiriage={kiriage}
                bakaze={bakaze}
                jikaze={jikaze}
                dora={dora}
                uradora={uradora}
              />
            )}

            {hand.length +
              chi.length +
              pon.length +
              (kan.length - parseInt(kan.length / 4)) +
              (ankan.length - parseInt(ankan.length / 4)) ===
              13 && (
              <Tenpai hand={hand} chi={chi} pon={pon} kan={kan} ankan={ankan} />
            )}

            <div className="handMain">
              <div className={hand.length ? "handPart" : ""}>
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

              <div className={chi.length ? "handPart" : ""}>
                {chi.map((tile, index) => (
                  <img
                    className="tile"
                    key={index}
                    src={`tiles/${tile}.png`}
                    draggable="false"
                    onClick={() => removeChi(tile, index)}
                  />
                ))}
              </div>

              <div className={pon.length ? "handPart" : ""}>
                {pon.map((tile, index) => (
                  <img
                    className="tile"
                    key={index}
                    src={`tiles/${tile}.png`}
                    draggable="false"
                    onClick={() => removePon(tile, index)}
                  />
                ))}
              </div>

              <div className={kan.length ? "handPart" : ""}>
                {kan.map((tile, index) => (
                  <img
                    className="tile"
                    key={index}
                    src={`tiles/${tile}.png`}
                    draggable="false"
                    onClick={() => removeKan(tile, index)}
                  />
                ))}
              </div>

              <div className={ankan ? "handPart" : ""}>
                {ankan.map((tile, index) => (
                  <img
                    className="tile"
                    key={index}
                    src={`tiles/${tile}.png`}
                    draggable="false"
                    onClick={() => removeAnkan(tile, index)}
                  />
                ))}
              </div>
            </div>

            <div>
              {hand.length +
                chi.length +
                pon.length +
                kan.length +
                ankan.length -
                parseInt(ankan.length / 4) >
                0 && (
                <button className="clearHand" onClick={() => clearHand()}>
                  Clear
                </button>
              )}

              <div>
                <button
                  onClick={() => changeCallType("chi")}
                  className={callType === "chi" ? "" : "off"}
                >
                  Chi
                </button>
                <button
                  onClick={() => changeCallType("pon")}
                  className={callType === "pon" ? "" : "off"}
                >
                  Pon
                </button>
                <button
                  onClick={() => changeCallType("kan")}
                  className={callType === "kan" ? "" : "off"}
                >
                  Kan
                </button>
                <button
                  onClick={() => changeCallType("ankan")}
                  className={callType === "ankan" ? "" : "off"}
                >
                  Ankan
                </button>
                <button
                  onClick={() => setAkaCall(!akaCall)}
                  className={akaCall ? "" : "off"}
                >
                  Call with akadora
                </button>
              </div>
            </div>

            <AddTiles
              hand={hand}
              setHand={(array) => setHand(array)}
              chi={chi}
              setChi={(array) => setChi(array)}
              pon={pon}
              setPon={(array) => setPon(array)}
              kan={kan}
              setKan={(array) => setKan(array)}
              ankan={ankan}
              setAnkan={(array) => setAnkan(array)}
              callType={callType}
              setCallType={(string) => setCallType(string)}
              akaCall={akaCall}
              tileCount={tileCount}
              setTileCount={(map) => setTileCount(map)}
              dora={dora}
              setDora={(array) => setDora(array)}
              uradora={uradora}
              setUradora={(array) => setUradora(array)}
              doraIndex={doraIndex}
              setDoraIndex={setDoraIndex}
              uradoraIndex={uradoraIndex}
              setUradoraIndex={setUradoraIndex}
            />
          </div>
        </div>

        <div className="rightMain">
          <div className="doraMain">
            <div>
              <Dora
                tileCount={tileCount}
                setTileCount={(map) => setTileCount(map)}
                dora={dora}
                setDora={(array) => setDora(array)}
                uradora={uradora}
                setUradora={(array) => setUradora(array)}
                doraIndex={doraIndex}
                setDoraIndex={setDoraIndex}
                uradoraIndex={uradoraIndex}
                setUradoraIndex={setUradoraIndex}
              />
            </div>
            <button className="doraClear" onClick={() => clearDora()}>
              Clear
            </button>
          </div>

          <ScoringOptions
            chi={chi}
            pon={pon}
            kan={kan}
            ankan={ankan}
            oya={oya}
            setOya={setOya}
            agari={agari}
            setAgari={setAgari}
            riichi={riichi}
            setRiichi={setRiichi}
            ippatsu={ippatsu}
            setIppatsu={setIppatsu}
            chankan={chankan}
            setChankan={setChankan}
            rinshan={rinshan}
            setRinshan={setRinshan}
            haitei={haitei}
            setHaitei={setHaitei}
            tenhou={tenhou}
            setTenhou={setTenhou}
            honba={honba}
            setHonba={setHonba}
            riichibou={riichibou}
            setRiichibou={setRiichibou}
            kiriage={kiriage}
            setKiriage={setKiriage}
            bakaze={bakaze}
            setBakaze={setBakaze}
            jikaze={jikaze}
            setJikaze={setJikaze}
          />
        </div>
      </div>
    </>
  );
}

export default Hand;
