import "./About.css";
function About() {
  return (
    <>
      <h1>About</h1>
      <hr />
      <div className="main">
        <h3>Riichi mahjong score calculator</h3>
        <p>
          Mahjong is a game with a lot of rules which can be overwhelming for
          newer players. One of the hardest parts to remember, even for more
          experienced players, is the scoring of a hand, specifically fu. This
          tool was made to make the process of scoring a hand easier and allows
          the user enters a hand. The tool will also show tenpai when 13 tiles
          are in the hand so it works as a tenpai finder as well which is useful
          for beginners and more complicated waits. The tool will show total
          score of a hand along with both dealer and non dealer payments. It
          also shows all the yaku included in the hand with their han amount as
          well as the fu of the hand. Users can also add dora, uradora, and
          akadora as well as adjust options that affect the score of the hand
          such as riichi and win type (ron / tsumo). The yaku page contains a
          list of all standard yaku along with a brief description of the yaku
          and example hands. The fu/scoring page has a brief overview on how to
          count a hands fu amount and calculate a score for a hand as well as
          useful scoring charts to memorize or streamline the scoring process.
          Thank you for visiting my website.
        </p>
      </div>

      <h1>Made With</h1>
      <hr />
      <div className="main">
        <h2>Go</h2>
        <p>
          Backend api written in go with chi router handles the most of the
          calculations of getting the score for the provided hand as well as
          checking tenpai for hands. The api has two endpoints being tenpai and
          score which respectively return a given hands tenpai tiles and hand
          score after calculating them.
        </p>
        <h2>Javascript React</h2>
        <p>
          This frontend website uses react to allow the user to interact with
          the backend by sending the hands to the go api endpoints then displays
          the given results back for the user in a nice and clean manner.
        </p>
      </div>

      <h1>References</h1>
      <hr />
      <div className="main">
        <p>
          Here are the main references I used when working on this tool. Riichi
          wiki is great for exact rules and I recommend taking a look as it has
          much more detail than I could write on this site
        </p>
        <a className="link" href="https://riichi.wiki/Main_Page">
          Riichi Wiki
        </a>
        <a className="link" href="https://mj-king.net/sozai/">
          Mahjong Oukoku (Tile Images)
        </a>
      </div>

      <h1>Contact</h1>
      <hr />
      <div className="main">
        <p>
          If you find any bugs / issues with the site or would like to get in
          contact with me, please let me know through either email or github.
        </p>

        <a
          className="link"
          href="https://github.com/KainoaGardner/webMahjongCalc"
        >
          Github
        </a>

        <a className="link" href="none">
          My Website (currently down)
        </a>

        <a className="link" href="mailto: kainoagardner123@gmail.com">
          kainoagardner123@gmail.com
        </a>
      </div>
    </>
  );
}
export default About;
