import "./Fu.css";
function Fu() {
  return (
    <>
      <h1>Fu</h1>
      <hr />

      <h2>Fu Total</h2>
      <div className="fuItem">
        <h3>Chiitoitsu always 25 fu</h3>
        <p>Else</p>
        <h3>All hands start with 20 fu </h3>
        <h3>After counting hand fu round up by 10</h3>
      </div>
      <div className="fuItem">
        <h3>Yakuhai Pair</h3>
        <h3>2 fu</h3>
        <p>If pair is a dragon or jikaze or bakaze 2 fu</p>
      </div>
      <div className="fuItem">
        <h3>Triplet</h3>
        <table className="tableTriple">
          <tr>
            <th></th>
            <th>Simples</th>
            <th>Terminals/Honors</th>
          </tr>
          <tr>
            <th>Open Triplet</th>
            <th>2 fu</th>
            <th>4 fu</th>
          </tr>
          <tr>
            <th>Closed Triplet</th>
            <th>4 fu</th>
            <th>8 fu</th>
          </tr>
          <tr>
            <th>Open Kan</th>
            <th>8 fu</th>
            <th>16 fu</th>
          </tr>
          <tr>
            <th>Ankan</th>
            <th>16 fu</th>
            <th>32 fu</th>
          </tr>
        </table>
        <p>Each triplet or kan gives fu shown by this chart</p>
      </div>

      <div className="fuItem">
        <h3>Wait</h3>
        <table className="tableTriple">
          <tr>
            <th>Ryanmen (two sided wait)</th>
            <th>0 fu</th>
          </tr>
          <tr>
            <th>Shanpon (dual pair wait)</th>
            <th>0* get fu for made triplet</th>
          </tr>
          <tr>
            <th>Kanchan (middle wait)</th>
            <th>2 fu</th>
          </tr>
          <tr>
            <th>Penchan (edge wait)</th>
            <th>2 fu</th>
          </tr>
          <tr>
            <th>Tanki (pair wait)</th>
            <th>2 fu</th>
          </tr>
        </table>
        <p>Fu added for type of wait</p>
      </div>

      <div className="fuItem">
        <h3>Winning Conditon</h3>
        <h4>Closed Ron</h4>
        <h4>2 fu</h4>
        <h4>Open or Closed Tsumo</h4>
        <h4>2 fu</h4>
        <p>
          Unless the hand has the yaku pinfu where it is then given 20 fu
          instead
        </p>
      </div>

      <div className="fuItem">
        <h3>Special</h3>
        <h4>Chiitoitsu</h4>
        <h4>25 fu</h4>
        <h4>Hand with no added fu</h4>
        <h4>30 fu</h4>
      </div>

      <h1>Scoring</h1>
      <hr />

      <h2>Basic Points</h2>
      <div className="fuItem">
        <h3>5 Han Mangan</h3>
        <h4>2000 Basic Points</h4>
        <hr />
        <h3>6,7 Han Haneman</h3>
        <h4>3000 Basic Points</h4>
        <hr />
        <h3>8,9,10 Han Baiman</h3>
        <h4>4000 Basic Points</h4>
        <hr />
        <h3>11,12 Han Sanbaiman</h3>
        <h4>6000 Basic Points</h4>
        <hr />
        <h3>13+ Kazoe Yakuman</h3>
        <h4>8000 Basic Points</h4>
        <hr />
        <h3>Yakuman</h3>
        <h4>8000 * Yakuman Amount = Basic Points</h4>
        <hr />

        <h3>&lt;5 Han</h3>
        <h4>fu * 2^(2 + han) = Basic Points</h4>
      </div>

      <h2>Score Payments</h2>
      <div className="fuItem">
        <h3>Payments</h3>
        <p>Payments based on if you are dealer or not and win type</p>
        <p>(oya payment / ko payment)</p>
        <table className="tablePayments">
          <tr>
            <th></th>
            <th>Ko (non dealer)</th>
            <th>Oya (dealer)</th>
          </tr>
          <tr>
            <th>Tsumo</th>
            <th>(2 * basic Points / 1 * Basic Points)</th>
            <th>2 * Basic Points</th>
          </tr>
          <tr>
            <th>Ron</th>
            <th>4 * Basic Points</th>
            <th>6 * Basic Points</th>
          </tr>
        </table>
      </div>

      <div className="fuItem">
        <h3>Honba</h3>
        <p>For each repeated dealer (honba) add 300 points to score</p>
        <h4>Ron</h4>
        <p>Player pays an extra 300 * honba</p>
        <h4>Tsumo</h4>
        <p>Each player pays an extra 100 * honba</p>
      </div>

      <h1>Scoring Charts</h1>
      <hr />
      <div className="fuMain">
        <h2>Ko (Non dealer)</h2>
        <table className="tableMain">
          <tr>
            <th></th>
            <th>20 Fu</th>
            <th>25 Fu</th>
            <th>30 Fu</th>
            <th>40 Fu</th>
            <th>50 Fu</th>
            <th>60 Fu</th>
            <th>70 Fu</th>
            <th>80 Fu</th>
            <th>90 Fu</th>
            <th>100 Fu</th>
            <th>110 Fu</th>
          </tr>
          <tr>
            <th>1 Han</th>
            <th></th>
            <th></th>
            <th>
              1000<p>(500/300)</p>
            </th>
            <th>
              1300<p>(700/400)</p>
            </th>
            <th>
              1600<p>(800/400)</p>
            </th>
            <th>
              2000<p>(1000/500)</p>
            </th>
            <th>
              2300<p>(1200/600)</p>
            </th>
            <th>
              2600<p>(1300/700)</p>
            </th>
            <th>
              2900<p>(1500/800)</p>
            </th>
            <th>
              3200<p>(1600/800)</p>
            </th>
            <th>
              3600<p>(1800/900)</p>
            </th>
          </tr>
          <tr>
            <th>2 Han</th>
            <th>
              1300<p>(700/400)</p>
            </th>
            <th>1600</th>
            <th>
              2000<p>(1000/500)</p>
            </th>
            <th>
              2600<p>(1300/700)</p>
            </th>
            <th>
              3200<p>(1600/800)</p>
            </th>
            <th>
              3900<p>(2000/1000)</p>
            </th>
            <th>
              4500<p>(2300/1200)</p>
            </th>
            <th>
              5200<p>(2600/1300)</p>
            </th>
            <th>
              5800<p>(2900/1500)</p>
            </th>
            <th>
              6400<p>(3200/1600)</p>
            </th>
            <th>
              7100<p>(3600/1800)</p>
            </th>
          </tr>
          <tr>
            <th>3 Han</th>
            <th>
              2600<p>(1300/700)</p>
            </th>
            <th>
              3200<p>(1600/800)</p>
            </th>
            <th>
              3900<p>(2000/1000)</p>
            </th>
            <th>
              5200<p>(2600/1300)</p>
            </th>
            <th>
              6400<p>(3200/1600)</p>
            </th>
            <th>
              7700<p>(3900/2000)</p>
            </th>

            <th colSpan="5">
              8000<p>(4000,2000)</p>
            </th>
          </tr>
          <tr>
            <th>4 Han</th>

            <th>
              5200<p>(2600/1300)</p>
            </th>
            <th>
              6400<p>(3200/1600)</p>
            </th>
            <th>
              7700<p>(3900/2000)</p>
            </th>
            <th colSpan="8">
              8000<p>(4000,2000)</p>
            </th>
          </tr>
          <tr>
            <th>5 Han</th>
            <th colSpan="11">
              8000<p>(4000,2000)</p>
            </th>
          </tr>
          <tr>
            <th>6,7 Han</th>
            <th colSpan="11">
              12000<p>(6000,3000)</p>
            </th>
          </tr>
          <tr>
            <th>8,9,10 Han</th>
            <th colSpan="11">
              16000<p>(8000,4000)</p>
            </th>
          </tr>
          <tr>
            <th>11,12 Han</th>
            <th colSpan="11">
              24000<p>(12000,6000)</p>
            </th>
          </tr>
          <tr>
            <th>13+ Han</th>
            <th colSpan="11">
              32000<p>(16000,8000)</p>
            </th>
          </tr>
        </table>

        <h2>Oya (Dealer)</h2>
        <table className="tableMain">
          <tr>
            <th></th>
            <th>20 Fu</th>
            <th>25 Fu</th>
            <th>30 Fu</th>
            <th>40 Fu</th>
            <th>50 Fu</th>
            <th>60 Fu</th>
            <th>70 Fu</th>
            <th>80 Fu</th>
            <th>90 Fu</th>
            <th>100 Fu</th>
            <th>110 Fu</th>
          </tr>
          <tr>
            <th>1 Han</th>
            <th></th>
            <th></th>
            <th>
              1500<p>(500)</p>
            </th>
            <th>
              2000<p>(700)</p>
            </th>
            <th>
              2400<p>(800)</p>
            </th>
            <th>
              2900<p>(1000)</p>
            </th>
            <th>
              3400<p>(1200)</p>
            </th>
            <th>
              3900<p>(1300)</p>
            </th>
            <th>
              4400<p>(1500)</p>
            </th>
            <th>
              4800<p>(1600)</p>
            </th>
            <th>
              5300<p>(1800)</p>
            </th>
          </tr>
          <tr>
            <th>2 Han</th>
            <th>
              2000<p>(700)</p>
            </th>
            <th>2400</th>
            <th>
              2900<p>(1000)</p>
            </th>
            <th>
              3900<p>(1300)</p>
            </th>
            <th>
              4800<p>(1600)</p>
            </th>
            <th>
              5800<p>(2000)</p>
            </th>
            <th>
              6800<p>(2300)</p>
            </th>
            <th>
              7700<p>(2600)</p>
            </th>
            <th>
              8700<p>(2900)</p>
            </th>
            <th>
              9600<p>(3200)</p>
            </th>
            <th>
              10600<p>(3600)</p>
            </th>
          </tr>
          <tr>
            <th>3 Han</th>
            <th>
              3900<p>(1300)</p>
            </th>
            <th>
              4800<p>(1600)</p>
            </th>
            <th>
              5800<p>(2000)</p>
            </th>
            <th>
              7700<p>(2600)</p>
            </th>
            <th>
              9600<p>(3200)</p>
            </th>
            <th>
              11600<p>(3900)</p>
            </th>

            <th colSpan="5">
              12000<p>(4000)</p>
            </th>
          </tr>
          <tr>
            <th>4 Han</th>

            <th>
              7700<p>(2600)</p>
            </th>
            <th>
              9600<p>(3200)</p>
            </th>
            <th>
              11600<p>(3900)</p>
            </th>
            <th colSpan="8">
              12000<p>(4000)</p>
            </th>
          </tr>
          <tr>
            <th>5 Han</th>
            <th colSpan="11">
              12000<p>(4000)</p>
            </th>
          </tr>
          <tr>
            <th>6,7 Han</th>
            <th colSpan="11">
              18000<p>(6000)</p>
            </th>
          </tr>
          <tr>
            <th>8,9,10 Han</th>
            <th colSpan="11">
              24000<p>(8000)</p>
            </th>
          </tr>
          <tr>
            <th>11,12 Han</th>
            <th colSpan="11">
              36000<p>(12000)</p>
            </th>
          </tr>
          <tr>
            <th>13+ Han</th>
            <th colSpan="11">
              48000<p>(16000)</p>
            </th>
          </tr>
        </table>
      </div>
    </>
  );
}
export default Fu;
