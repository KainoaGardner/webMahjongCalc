import { useState, useEffect } from "react";
import "./ScoringOptions.css";

function ScoringOptions() {
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

  function changeRiichi() {
    // if (riichi === "Dama" && OPENHAND) {
    //   alert("Cannot call riichi with an open hand");
    //   return;
    // }
    switch (riichi) {
      case "None":
        setRiichi("Riichi");
        break;
      case "Riichi":
        setRiichi("Wriichi");
        break;

      default:
        setRiichi("None");
    }
  }

  function changeIppatsu() {
    if (riichi === "None") {
      alert("Must have riichi or wriichi for ippatsu");
      return;
    }
    setIppatsu(!ippatsu);
  }

  function changeChankan() {
    if (!agari) {
      alert("Must win by ron for chankan");
      return;
    }
    setChankan(!chankan);
  }

  function changeRinshan() {
    if (agari) {
      alert("Must win by tsmuo for rinshan");
      return;
    }
    // if (!rinshan && !KAN) {
    //   alert("Must have at least 1 kan for rinshan");
    //   return;
    // }

    setRinshan(!rinshan);
  }

  function changeHaitei() {
    if (haitei === "None" && agari) {
      setHaitei("Houtei");
    } else if (haitei === "None" && !agari) {
      setHaitei("Haitei");
    } else {
      setHaitei("None");
    }
  }

  function changeTenhou() {
    if (agari) {
      alert("Cant win Tenhou or chiihou with ron");
      return;
    }
    if (tenhou === "None" && oya) {
      setTenhou("Tenhou");
    } else if (tenhou === "None" && !oya) {
      setTenhou("Chiihou");
    } else {
      setTenhou("None");
    }
  }

  function decreaseHonba() {
    if (honba === 0) {
      alert("Cant have negative honba");
      return;
    }
    setHonba(honba - 1);
  }

  function decreaseRiichibou() {
    if (riichibou === 0) {
      alert("Cant have negative riichibou");
      return;
    }
    setRiichibou(riichibou - 1);
  }

  function changeBakaze() {
    switch (bakaze) {
      case "H1":
        setBakaze("H2");
        break;
      case "H2":
        setBakaze("H3");
        break;
      case "H3":
        setBakaze("H4");
        break;
      case "H4":
        setBakaze("H1");
        break;
    }
  }

  function changeJikaze() {
    switch (jikaze) {
      case "H1":
        setJikaze("H2");
        break;
      case "H2":
        setJikaze("H3");
        break;
      case "H3":
        setJikaze("H4");
        break;
      case "H4":
        setJikaze("H1");
        break;
    }
  }

  useEffect(() => {
    setHaitei("None");
    setTenhou("None");
    setRinshan(false);
    setChankan(false);
  }, [agari]);

  useEffect(() => {
    setTenhou("None");
  }, [oya]);

  useEffect(() => {
    if (riichi === "None") {
      setIppatsu(false);
    }
  }, [riichi]);

  return (
    <>
      <div>
        <button onClick={() => setOya(!oya)}>{oya ? "Oya" : "Ko"}</button>
        <button onClick={() => setAgari(!agari)}>
          {agari ? "Ron" : "Tsumo"}
        </button>
        <button
          onClick={() => changeRiichi()}
          className={riichi === "None" ? "off" : ""}
        >
          {riichi === "None" ? "Riichi" : riichi}
        </button>
        <button
          onClick={() => changeIppatsu()}
          className={ippatsu ? "" : "off"}
        >
          Ippatsu
        </button>
        <button
          onClick={() => changeChankan()}
          className={chankan ? "" : "off"}
        >
          Chankan
        </button>
        <button
          onClick={() => changeRinshan()}
          className={rinshan ? "" : "off"}
        >
          Rinshan
        </button>
        <button
          onClick={() => changeHaitei()}
          className={haitei === "None" ? "off" : ""}
        >
          {(() => {
            if (haitei === "None" && agari) {
              return "Houtei";
            } else if (haitei === "None" && !agari) {
              return "Haitei";
            } else {
              return haitei;
            }
          })()}
        </button>
        <button
          onClick={() => changeTenhou()}
          className={tenhou === "None" ? "off" : ""}
        >
          {(() => {
            if (tenhou === "None" && oya) {
              return "Tenhou";
            } else if (tenhou === "None" && !oya) {
              return "Chiihou";
            } else {
              return tenhou;
            }
          })()}
        </button>
        <button
          onClick={() => setKiriage(!kiriage)}
          className={kiriage ? "" : "off"}
        >
          Kiriage
        </button>

        <div>
          <div>
            <h2>Honba</h2>
            <button onClick={() => decreaseHonba()}>-</button>
            {honba}
            <button onClick={() => setHonba(honba + 1)}>+</button>
          </div>

          <div>
            <h2>Riichibou</h2>
            <button onClick={() => decreaseRiichibou()}>-</button>
            {riichibou}
            <button onClick={() => setRiichibou(riichibou + 1)}>+</button>
          </div>
        </div>

        <div>
          <div>
            <h2>Bakaze</h2>
            <img
              className="tile"
              src={`tiles/${bakaze}.png`}
              draggable="false"
              onClick={() => changeBakaze()}
            />
          </div>
          <div>
            <h2>Jikaze</h2>
            <img
              className="tile"
              src={`tiles/${jikaze}.png`}
              draggable="false"
              onClick={() => changeJikaze()}
            />
          </div>
        </div>
      </div>
    </>
  );
}

export default ScoringOptions;
