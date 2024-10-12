import "./Hand.css";
import "./Results.css";
import { useState, useEffect } from "react";

function Results({
  hand,
  chi,
  pon,
  kan,
  ankan,
  dora,
  uradora,
  oya,
  agari,
  riichi,
  ippatsu,
  chankan,
  rinshan,
  haitei,
  tenhou,
  honba,
  riichibou,
  kiriage,
  bakaze,
  jikaze,
}) {
  const [menzenResult, setMenzenResult] = useState([]);
  const [chiResult, setChiResult] = useState([]);
  const [ponResult, setPonResult] = useState([]);
  const [kanResult, setKanResult] = useState([]);
  const [ankanResult, setAnkanResult] = useState([]);

  const [score, setScore] = useState(0);
  const [scoreType, setScoreType] = useState("");
  const [koPayment, setKoPayment] = useState(0);
  const [oyaPayment, setOyaPayment] = useState(0);
  const [fu, setFu] = useState(0);
  const [fuList, setFuList] = useState([]);
  const [han, setHan] = useState(0);
  const [yakuList, setYakuList] = useState([]);
  const [yakuman, setYakuman] = useState(0);
  const [yakumanList, setYakumanList] = useState([]);

  const API_URL =
    import.meta.env.VITE_API_HOST +
    ":" +
    import.meta.env.VITE_API_PORT +
    "/api/" +
    import.meta.env.VITE_API_VER +
    "/";

  function getFixedAnkan() {
    const newAnkan = [];
    for (let i = 0; i < ankan.length; i++) {
      if (i % 4 == 0) {
        newAnkan[i] = ankan[i + 1];
      } else if (i % 4 === 3) {
        newAnkan[i] = ankan[i - 2];
      } else {
        newAnkan[i] = ankan[i];
      }
    }
    return newAnkan;
  }

  function getDora(doraIndicator) {
    const dora = [];
    for (let i = 0; i < doraIndicator.length; i++) {
      if (doraIndicator[i] === "B0") {
        continue;
      }
      const tile = doraIndicator[i];
      let tileNumber = parseInt(tile[1]) + 1;
      if (tile[0] === "H") {
        if (tileNumber === 5) {
          tileNumber = 1;
        } else if (tileNumber === 8) {
          tileNumber = 5;
        }
      } else {
        if (tileNumber > 9) {
          tileNumber = 1;
        }
      }

      tileNumber = tileNumber.toString();
      let newTile = tile[0] + tileNumber;
      if (tile.length > 2) {
        newTile += tile[2];
      }

      dora.push(newTile);
    }
    return dora;
  }

  const handParts = {
    Menzen: hand,
    Chi: chi,
    Pon: pon,
    Kan: kan,
    Ankan: getFixedAnkan(),
  };

  const scoringParts = {
    Dora: getDora(dora),
    Uradora: getDora(uradora),
    Oya: oya,
    Bakaze: bakaze,
    Jikaze: jikaze,
    Ron: agari,
    Tsumo: !agari,
    Riichi: riichi === "Riichi",
    Wriichi: riichi === "Wriichi",
    Ippatsu: ippatsu,
    Chankan: chankan,
    Rinshan: rinshan,
    Haitei: haitei === "Haitei",
    Houtei: haitei === "Houtei",
    Tenhou: tenhou === "Tenhou",
    Chiihou: tenhou === "Chiihou",
    Honba: honba,
    RiichiBou: riichibou,
    Kiriage: kiriage,
  };

  const scoreHandPost = {
    Hand: handParts,
    ScoringParts: scoringParts,
  };

  const scorePost = async () => {
    console.log(scoreHandPost);
    try {
      const response = await fetch("http://" + API_URL + "points", {
        method: "POST",
        headers: {
          "Content-Type": "application/json; charset=utf-8",
        },
        body: JSON.stringify(scoreHandPost),
      });

      if (response.ok) {
        const data = await response.json();
        console.log(data);
        changeScore(data);
      }
    } catch (error) {
      console.log(error);
      setScore();
    }
  };

  function changeScore(data) {
    setScore(data.handScore.score);
    setScoreType(data.handScore.scoreType);

    setKoPayment(data.handScore.koPayment);
    setOyaPayment(data.handScore.oyaPayment);

    setFu(data.handScore.fu);
    if (data.handScore.fuList) {
      setFuList(data.handScore.fuList);
    }

    setHan(data.handScore.han);
    if (data.handScore.yakuList) {
      setYakuList(data.handScore.yakuList);
    }

    setYakuman(data.handScore.yakuman);
    if (data.handScore.yakumanList) {
      setYakumanList(data.handScore.yakumanList);
    }

    if (data.hand.menzen) {
      setMenzenResult(data.hand.menzen);
    }
    if (data.hand.chi) {
      setChiResult(data.hand.chi);
    }
    if (data.hand.pon) {
      setPonResult(data.hand.pon);
    }
    if (data.hand.kan) {
      setKanResult(data.hand.kan);
    }
    if (data.hand.ankan) {
      setAnkanResult(data.hand.ankan);
    }
  }

  useEffect(() => {
    scorePost();
  }, [
    agari,
    oya,
    riichi,
    ippatsu,
    chankan,
    rinshan,
    haitei,
    tenhou,
    honba,
    riichibou,
    kiriage,
    bakaze,
    jikaze,
    dora,
    uradora,
  ]);

  return (
    <>
      <div className="result">
        <h2 className="point">{scoreType}</h2>
        <h2 className="point">{score !== 0 ? score : ""} </h2>

        <div className="payment">
          <div>
            {!agari ? (
              <div className="scoreCompMain">
                {!oya ? (
                  <p className="scoreComp">Oya Payment : {oyaPayment}</p>
                ) : (
                  <></>
                )}
                <p className="scoreComp">Ko Payment : {koPayment}</p>
              </div>
            ) : (
              <></>
            )}
          </div>
        </div>
      </div>

      <div className="resultMain">
        <div className={han ? "scoreBlock" : ""}>
          {han ? (
            <div>
              <h3 className="scoreCompTitle">Han: {han}</h3>
              <hr />
              <div className="scoreCompMain">
                {yakuList.map((yakuComp, index) => (
                  <p key={index} className="scoreComp">
                    {yakuComp.title} : {yakuComp.han}
                  </p>
                ))}
              </div>
            </div>
          ) : (
            <></>
          )}
        </div>

        <div className={fu ? "scoreBlock" : ""}>
          {fu ? (
            <div>
              <h3 className="scoreCompTitle">Fu: {fu}</h3>
              <hr />
              <div className="scoreCompMain">
                {fuList.map((fuComp, index) => (
                  <p key={index} className="scoreComp">
                    {fuComp.title} : {fuComp.fu}
                  </p>
                ))}
              </div>
            </div>
          ) : (
            <></>
          )}
        </div>

        <div className={yakuman ? "scoreBlock" : ""}>
          {yakuman ? (
            <div>
              <h3 className="scoreCompTitle">Yakuman: {yakuman}</h3>
              <hr />
              <div className="scoreCompMain">
                {yakumanList.map((yakumanComp, index) => (
                  <p key={index} className="scoreComp">
                    {yakumanComp.title} : {yakumanComp.yakuman}
                  </p>
                ))}
              </div>
            </div>
          ) : (
            <></>
          )}
        </div>
      </div>

      {menzenResult.length === 0 ? <h2>Invalid Hand</h2> : <></>}
      <div className="resultHandMain">
        <div className={menzenResult ? "resultHandPart" : ""}>
          {menzenResult.map((block, blockIndex) => (
            <div key={blockIndex} className="resultHandBlock">
              {block.map((tile, index) => (
                <img
                  className="tile"
                  key={index}
                  src={`tiles/${tile}.png`}
                  draggable="false"
                />
              ))}
            </div>
          ))}
        </div>
        <div className={chiResult ? "resultHandPart" : ""}>
          {chiResult.map((block, blockIndex) => (
            <div key={blockIndex} className="resultHandBlock">
              {block.map((tile, index) => (
                <img
                  className="tile"
                  key={index}
                  src={`tiles/${tile}.png`}
                  draggable="false"
                />
              ))}
            </div>
          ))}
        </div>
        <div className={chiResult ? "resultHandPart" : ""}>
          {ponResult.map((block, blockIndex) => (
            <div key={blockIndex} className="resultHandBlock">
              {block.map((tile, index) => (
                <img
                  className="tile"
                  key={index}
                  src={`tiles/${tile}.png`}
                  draggable="false"
                />
              ))}
            </div>
          ))}
        </div>
        <div className={chiResult ? "resultHandPart" : ""}>
          {kanResult.map((block, blockIndex) => (
            <div key={blockIndex} className="resultHandBlock">
              {block.map((tile, index) => (
                <img
                  className="tile"
                  key={index}
                  src={`tiles/${tile}.png`}
                  draggable="false"
                />
              ))}
            </div>
          ))}
        </div>
        <div className={chiResult ? "resultHandPart" : ""}>
          {ankanResult.map((block, blockIndex) => (
            <div key={blockIndex} className="resultHandBlock">
              {block.map((tile, index) => (
                <img
                  className="tile"
                  key={index}
                  src={`tiles/${index % 4 === 0 || index % 4 === 3 ? "B0" : tile}.png`}
                  draggable="false"
                />
              ))}
            </div>
          ))}
        </div>
      </div>
      <hr />
    </>
  );
}

export default Results;
