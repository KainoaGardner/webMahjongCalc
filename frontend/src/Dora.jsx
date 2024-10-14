import "./Hand.css";
import "./Dora.css";

function Dora({
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
  function changeDora(tile, index) {
    if (index === doraIndex) {
      setDoraIndex(-1);
      return;
    }
    if (tile !== "B0") {
      dora[index] = "B0";
      setDora((d) => [...d]);
      setTileCount(new Map(tileCount).set(tile, tileCount.get(tile) - 1));
    } else {
      setDoraIndex(index);
      setUradoraIndex(-1);
    }
  }

  function changeUradora(tile, index) {
    if (index === uradoraIndex) {
      setUradoraIndex(-1);
      return;
    }
    if (tile !== "B0") {
      uradora[index] = "B0";
      setUradora((d) => [...d]);
      setTileCount(new Map(tileCount).set(tile, tileCount.get(tile) - 1));
    } else {
      setUradoraIndex(index);
      setDoraIndex(-1);
    }
  }

  return (
    <>
      <div>
        <h2 className="doraTitle">Dora Indicator</h2>
        {dora.map((tile, index) => (
          <img
            className={doraIndex === index ? "selected tile" : "tile"}
            key={index}
            src={`tiles/${tile}.png`}
            draggable="false"
            onClick={() => changeDora(tile, index)}
          />
        ))}
      </div>

      <div>
        <h2 className="doraTitle">Uradora Indicator</h2>
        {uradora.map((tile, index) => (
          <img
            className={uradoraIndex === index ? "selected tile" : "tile"}
            key={index}
            src={`tiles/${tile}.png`}
            draggable="false"
            onClick={() => changeUradora(tile, index)}
          />
        ))}
      </div>
    </>
  );
}

export default Dora;
