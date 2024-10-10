import "./Hand.css";
import { useState, useEffect } from "react";

function Tenpai({ hand, chi, pon, kan, ankan }) {
  const [tenpai, setTenpai] = useState([]);

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

  const tenpaiPost = async () => {
    try {
      const response = await fetch("http://" + API_URL + "tenpai", {
        method: "POST",
        headers: {
          "Content-Type": "application/json; charset=utf-8",
        },
        body: JSON.stringify(handParts),
      });

      if (response.ok) {
        const data = await response.json();
        changeTenpai(data);
      }
    } catch (error) {
      console.log(error);
    }
  };

  function changeTenpai(data) {
    data.sort();
    setTenpai(data);
  }

  useEffect(() => {
    tenpaiPost();
  }, []);

  return (
    <>
      <h2>Tenpai</h2>
      <div>
        {tenpai.map((tile, index) => (
          <img
            className="tile"
            key={index}
            src={`tiles/${tile}.png`}
            draggable="false"
          />
        ))}
      </div>
    </>
  );
}

export default Tenpai;
