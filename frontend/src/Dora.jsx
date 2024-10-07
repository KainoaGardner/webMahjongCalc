import { useState } from "react";

function Dora() {
  const [dora, setDora] = useState(["M1", "B0", "B0", "B0"]);
  const [uradora, setUradora] = useState(["B0", "B0", "B0", "B0"]);

  function removeDoraTile(tile, index) {
    if (tile !== "B0") {
      dora[index] = "B0";
      setDora((d) => [...d]);
    }
  }

  return (
    <>
      <div>
        {dora.map((tile, index) => (
          <img
            className="tile"
            key={index}
            src={`tiles/${tile}.png`}
            draggable="false"
            onClick={() => removeDoraTile(tile, index)}
          />
        ))}
      </div>

      <div>
        {uradora.map((tile, index) => (
          <img
            className="tile"
            key={index}
            src={`tiles/${tile}.png`}
            draggable="false"
            onClick={() => removeDoraTile(tile, index)}
          />
        ))}
      </div>
    </>
  );
}

export default Dora;
