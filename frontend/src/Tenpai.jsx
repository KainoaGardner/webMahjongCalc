import "./Tenpai.css";
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

  const handParts = {
    Menzen: hand,
    Chi: chi,
    Pon: pon,
    Kan: kan,
    Ankan: getFixedAnkan(),
  };

  const tenpaiPost = async () => {
    try {
      console.log(handParts);
      const response = await fetch("http://" + API_URL + "tenpai", {
        method: "POST",
        headers: {
          "Content-Type": "application/json; charset=utf-8",
        },
        body: JSON.stringify(handParts),
      });

      if (response.ok) {
        const data = await response.json();
        console.log(data);
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
      {tenpai.length !== 0 ? (
        <>
          {/* <h2>Tenpai</h2> */}
          <div className="tenpaiMain">
            {tenpai.map((tile, index) => (
              <img
                className="tenpaiTile"
                key={index}
                src={`tiles/${tile}.png`}
                draggable="false"
              />
            ))}
          </div>
        </>
      ) : (
        <h2>Noten</h2>
      )}
    </>
  );
}

export default Tenpai;
