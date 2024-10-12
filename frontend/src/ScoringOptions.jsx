import { useEffect } from "react";
import "./ScoringOptions.css";

function ScoringOptions({
  chi,
  pon,
  kan,
  ankan,
  oya,
  setOya,
  agari,
  setAgari,
  riichi,
  setRiichi,
  ippatsu,
  setIppatsu,
  chankan,
  setChankan,
  rinshan,
  setRinshan,
  haitei,
  setHaitei,
  tenhou,
  setTenhou,
  honba,
  setHonba,
  riichibou,
  setRiichibou,
  kiriage,
  setKiriage,
  bakaze,
  setBakaze,
  jikaze,
  setJikaze,
}) {
  function changeRiichi() {
    if (chi.length + pon.length + kan.length > 0) {
      setRiichi("None");
      return;
    }
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
      setIppatsu(false);
      return;
    }
    setIppatsu(!ippatsu);
  }

  function changeChankan() {
    if (!agari) {
      setChankan(false);
      return;
    }
    setChankan(!chankan);
  }

  function changeRinshan() {
    if (agari) {
      setRinshan(false);
      return;
    }
    if (!rinshan && kan.length + ankan.length == 0) {
      setRinshan(false);
      return;
    }

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
    if (agari || chi.length + pon.length + kan.length + ankan.length > 0) {
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
      return;
    }
    setHonba(honba - 1);
  }

  function decreaseRiichibou() {
    if (riichibou === 0) {
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

  function checkValidRiichi() {
    if (chi.length + pon.length + kan.length > 0) {
      return riichi === "None" ? "off notValid" : "notValid";
    }
    return riichi === "None" ? "off" : "on";
  }

  function checkValidIppatsu() {
    if (riichi === "None") {
      return ippatsu ? "notValid" : "off notValid";
    }
    return ippatsu ? "on" : "off";
  }

  function checkValidChankan() {
    if (!agari) {
      return chankan ? "notValid" : "off notValid";
    }
    return chankan ? "on" : "off";
  }

  function checkValidRinshan() {
    if (kan.length + ankan.length === 0 || agari) {
      return rinshan ? "notValid" : "off notValid";
    }

    return rinshan ? "on" : "off";
  }

  function checkValidTenhou() {
    if (agari || chi.length + pon.length + kan.length + ankan.length > 0) {
      return tenhou === "None" ? "off notValid" : "notValid";
    }

    return tenhou === "None" ? "off" : "on";
  }

  useEffect(() => {
    setRiichi("None");
  }, [chi, pon, kan]);

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
      <div className="kaze">
        <div className="kazePart">
          <h2 className="kazeTitle">Bakaze</h2>
          <img
            className="tile"
            src={`tiles/${bakaze}.png`}
            draggable="false"
            onClick={() => changeBakaze()}
          />
        </div>
        <div className="kazePart">
          <h2 className="kazeTitle">Jikaze</h2>
          <img
            className="tile"
            src={`tiles/${jikaze}.png`}
            draggable="false"
            onClick={() => changeJikaze()}
          />
        </div>
      </div>

      <div className="scoringOptions">
        <button className="on" onClick={() => setOya(!oya)}>
          {oya ? "Oya" : "Ko"}
        </button>
        <button className="on" onClick={() => setAgari(!agari)}>
          {agari ? "Ron" : "Tsumo"}
        </button>
        <button onClick={() => changeRiichi()} className={checkValidRiichi()}>
          {riichi === "None" ? "Riichi" : riichi}
        </button>
        <button onClick={() => changeIppatsu()} className={checkValidIppatsu()}>
          Ippatsu
        </button>
        <button onClick={() => changeChankan()} className={checkValidChankan()}>
          Chankan
        </button>
        <button onClick={() => changeRinshan()} className={checkValidRinshan()}>
          Rinshan
        </button>
        <button
          onClick={() => changeHaitei()}
          className={haitei === "None" ? "off" : "on"}
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
        <button onClick={() => changeTenhou()} className={checkValidTenhou()}>
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
          className={kiriage ? "on" : "off"}
        >
          Kiriage
        </button>

        <div>
          <div>
            <h2>Honba</h2>
            <button
              onClick={() => decreaseHonba()}
              className={honba === 0 ? "notValid" : ""}
            >
              -
            </button>
            {honba}
            <button onClick={() => setHonba(honba + 1)}>+</button>
          </div>

          <div>
            <h2>Riichibou</h2>
            <button
              onClick={() => decreaseRiichibou()}
              className={riichibou === 0 ? "notValid" : ""}
            >
              -
            </button>
            {riichibou}
            <button onClick={() => setRiichibou(riichibou + 1)}>+</button>
          </div>
        </div>
      </div>
    </>
  );
}

export default ScoringOptions;
