import "./Hand.css";
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

  const handParts = {
    Menzen: hand,
    Chi: chi,
    Pon: pon,
    Kan: kan,
    Ankan: ankan,
  };

  const scoringParts = {
    Dora: dora.filter((d) => d !== "B0"),
    Uradora: uradora.filter((d) => d !== "B0"),
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
  }, []);

  return (
    <>
      <div>
        <h2>Score</h2>
        <p>{score !== 0 ? score : ""}</p>
        <p>{scoreType}</p>

        {!agari ? (
          <div>
            {!oya ? <p>Oya Payment : {oyaPayment}</p> : <></>}
            <p>Ko Payment : {koPayment}</p>
          </div>
        ) : (
          <></>
        )}

        {han ? (
          <div>
            <h3>Yaku</h3>
            {han}
            <div>
              {yakuList.map((yakuComp, index) => (
                <p key={index}>
                  {yakuComp.title} : {yakuComp.han}
                </p>
              ))}
            </div>
          </div>
        ) : (
          <></>
        )}

        {yakuman ? (
          <div>
            <h3>Yakuman</h3>
            {yakuman}
            <div>
              {yakumanList.map((yakumanComp, index) => (
                <p key={index}>
                  {yakumanComp.title} : {yakumanComp.yakuman}
                </p>
              ))}
            </div>
          </div>
        ) : (
          <></>
        )}
      </div>

      <div>
        {menzenResult.map((block, blockIndex) => (
          <div key={blockIndex}>
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
      <div>
        {chiResult.map((block, blockIndex) => (
          <div key={blockIndex}>
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
      <div>
        {ponResult.map((block, blockIndex) => (
          <div key={blockIndex}>
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
      <div>
        {kanResult.map((block, blockIndex) => (
          <div key={blockIndex}>
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
      <div>
        {ankanResult.map((block, blockIndex) => (
          <div key={blockIndex}>
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
    </>
  );
}

export default Results;
