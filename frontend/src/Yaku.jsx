import "./Yaku.css";
function Yaku() {
  return (
    <>
      <h1>Yaku</h1>
      <hr />
      <div className="yakuList">
        <h2 className="hanTitle">1 Han</h2>
        <hr />
        <div className="yakuItem">
          <h2>Menzen Tsumo</h2>
          <hr />
          <h3>Closed: 1</h3>
          <h3>Open: 0</h3>
          <p>Draw your winning tile yourself with a closed hand</p>
        </div>

        <div className="yakuItem">
          <h2>Riichi</h2>
          <hr />
          <h3>Closed: 1</h3>
          <h3>Open: 0</h3>
          <p>
            Place a 1000 point bet while in tenpai with a closed hand to call
            riichi. Players in riichi are unable to change their hands after
            calling riichi
          </p>
        </div>

        <div className="yakuItem">
          <h2>Ippatsu</h2>
          <hr />
          <h3>Closed: 1</h3>
          <h3>Open: 0</h3>
          <p>
            Win within the next turn after calling riichi. Any calls interrupt
            ippatus
          </p>
        </div>

        <div className="yakuItem">
          <h2>Haitei</h2>
          <hr />
          <h3>Closed: 1</h3>
          <h3>Open: 1</h3>
          <p>Win by drawing the very last tile</p>
        </div>

        <div className="yakuItem">
          <h2>Houtei</h2>
          <hr />
          <h3>Closed: 1</h3>
          <h3>Open: 1</h3>
          <p>Win by the very last discarded tile</p>
        </div>

        <div className="yakuItem">
          <h2>Rinshan kaihou</h2>
          <hr />
          <h3>Closed: 1</h3>
          <h3>Open: 1</h3>
          <p>Win on the tile drawn from the dead wall after calling a kan</p>
        </div>

        <div className="yakuItem">
          <h2>Chankan</h2>
          <hr />
          <h3>Closed: 1</h3>
          <h3>Open: 1</h3>
          <p>
            Win off an opponent's tile when they add to an open pon to make it a
            kan
          </p>
        </div>

        <div className="yakuItem">
          <h2>Tanyao</h2>
          <hr />
          <h3>Closed: 1</h3>
          <h3>Open: 1</h3>
          <p>
            Win with only simple tiles (2-8) and no honor or terminal tiles (1
            or 9)
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M6",
              "M7",
              "M8",
              "S2",
              "S3",
              "S4",
              "S5",
              "S5",
              "S5",
              "P4",
              "P5",
              "P6",
              "P8",
              "P8",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Pinfu</h2>
          <hr />
          <h3>Closed: 1</h3>
          <h3>Open: 0</h3>
          <p>
            Win with only sequences and a non yakuhai pair. The wait must be two
            sided
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M7",
              "M8",
              "M9",
              "S2",
              "S3",
              "S4",
              "S5",
              "S6",
              "S7",
              "P4",
              "P5",
              "P6",
              "P8",
              "P8",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Iipeikou</h2>
          <hr />
          <h3>Closed: 1</h3>
          <h3>Open: 0</h3>
          <p>Have two of the same sequence</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "P3",
              "P3",
              "P4",
              "P4",
              "P5",
              "P5",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Yakuhai</h2>
          <hr />
          <h3>Closed: 1</h3>
          <h3>Open: 1</h3>
          <p>
            Have a triplet of any of the dragon tiles or the round or seat wind
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "H7",
              "H7",
              "H7",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
            ]}
          />
        </div>

        <h2 className="hanTitle">2 Han</h2>
        <hr />

        <div className="yakuItem">
          <h2>Double Riichi</h2>
          <hr />
          <h3>Closed: 2</h3>
          <h3>Open: 0</h3>
          <p>Call riichi on your very first discard</p>
        </div>

        <div className="yakuItem">
          <h2>Chiitoitsu</h2>
          <hr />
          <h3>Closed: 2</h3>
          <h3>Open: 0</h3>
          <p>
            Special hand won with 7 pairs instead of the 4 groups and a pair
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M1",
              "M1",
              "M8",
              "M8",
              "S3",
              "S3",
              "P5",
              "P5",
              "P9",
              "P9",
              "H1",
              "H1",
              "H5",
              "H5",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Sanshoku doujun</h2>
          <hr />
          <h3>Closed: 2</h3>
          <h3>Open: 1</h3>
          <p>Win with the same sequence in each suit</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M3",
              "M4",
              "M5",
              "S3",
              "S4",
              "S5",
              "P3",
              "P4",
              "P5",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Ittsu</h2>
          <hr />
          <h3>Closed: 2</h3>
          <h3>Open: 1</h3>
          <p>Win with the sequences 123 456 789 of the same suit</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "P1",
              "P2",
              "P3",
              "P4",
              "P5",
              "P6",
              "P7",
              "P8",
              "P9",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Chanta</h2>
          <hr />
          <h3>Closed: 2</h3>
          <h3>Open: 1</h3>
          <p>
            Win where all groups and pair have at least one terminal or honor
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M1",
              "M2",
              "M3",
              "S7",
              "S8",
              "S9",
              "P9",
              "P9",
              "P9",
              "H4",
              "H4",
              "H4",
              "H6",
              "H6",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Toitoi</h2>
          <hr />
          <h3>Closed: 2</h3>
          <h3>Open: 2</h3>
          <p>Win with all triplets (or kans) and a pair</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M3",
              "M3",
              "M3",
              "S1",
              "S1",
              "S1",
              "P7",
              "P7",
              "P9",
              "P9",
              "P9",
              "H5",
              "H5",
              "H5",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Shousangen</h2>
          <hr />
          <h3>Closed: 2</h3>
          <h3>Open: 2</h3>
          <p>
            Win with two triplets of the dragons and a pair of the last dragon.
            Is always 4 han when including the two yakuhai triplets
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "H5",
              "H5",
              "H5",
              "H6",
              "H6",
              "H6",
              "H7",
              "H7",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Sanankou</h2>
          <hr />
          <h3>Closed: 2</h3>
          <h3>Open: 2</h3>
          <p>
            Win with three closed triplets. The rest of your hand can be open as
            long as the triplets are closed.
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M5",
              "M5",
              "M5",
              "S9",
              "S9",
              "S9",
              "H2",
              "H2",
              "H2",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Honroutou</h2>
          <hr />
          <h3>Closed: 2</h3>
          <h3>Open: 2</h3>
          <p>
            Win with only terminal and honor tiles. Can be considered as 4 han
            since Honroutou will always also be toitoi
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M9",
              "M9",
              "M9",
              "S1",
              "S1",
              "S1",
              "S9",
              "S9",
              "S9",
              "P1",
              "P1",
              "P1",
              "H3",
              "H3",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Sanshoku doukou</h2>
          <hr />
          <h3>Closed: 2</h3>
          <h3>Open: 2</h3>
          <p>Win with the same triplet in each suit</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M6",
              "M6",
              "M6",
              "S6",
              "S6",
              "S6",
              "P6",
              "P6",
              "P6",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Sankantsu</h2>
          <hr />
          <h3>Closed: 2</h3>
          <h3>Open: 2</h3>
          <p>Win with three kans</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M1",
              "M1",
              "M1",
              "M1",
              "S4",
              "S4",
              "S4",
              "S4",
              "H1",
              "H1",
              "H1",
              "H1",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
            ]}
          />
        </div>

        <h2 className="hanTitle">3 Han</h2>
        <hr />

        <div className="yakuItem">
          <h2>Ryanpeikou</h2>
          <hr />
          <h3>Closed: 3</h3>
          <h3>Open: 0</h3>
          <p>
            Win with two sets of iipeikou (two of the same sequence). Does not
            combine with Chiitoitsu
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M6",
              "M6",
              "M7",
              "M7",
              "M8",
              "M8",
              "S2",
              "S2",
              "S3",
              "S3",
              "S4",
              "S4",
              "H5",
              "H5",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Honitsu</h2>
          <hr />
          <h3>Closed: 3</h3>
          <h3>Open: 2</h3>
          <p>Win with all of one suit and honors</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "S1",
              "S2",
              "S3",
              "S5",
              "S5",
              "S5",
              "S6",
              "S7",
              "S8",
              "H1",
              "H1",
              "H1",
              "H5",
              "H5",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Junchan</h2>
          <hr />
          <h3>Closed: 3</h3>
          <h3>Open: 2</h3>
          <p>Win with only groups that contain a terminal tile</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M1",
              "M2",
              "M3",
              "S1",
              "S1",
              "S1",
              "S7",
              "S8",
              "S9",
              "P1",
              "P1",
              "P7",
              "P8",
              "P9",
            ]}
          />
        </div>

        <h2 className="hanTitle">6 Han</h2>
        <hr />

        <div className="yakuItem">
          <h2>Chinitsu</h2>
          <hr />
          <h3>Closed: 6</h3>
          <h3>Open: 5</h3>
          <p>Win with only tiles of a single suit</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "P1",
              "P2",
              "P3",
              "P3",
              "P3",
              "P4",
              "P5",
              "P6",
              "P7",
              "P7",
              "P7",
              "P7",
              "P8",
              "P9",
            ]}
          />
        </div>

        <h2 className="hanTitle">Yakuman</h2>
        <hr />

        <div className="yakuItem">
          <h2>Kazoe yakuman</h2>
          <hr />
          <h3>Closed: Yakuman</h3>
          <h3>Open: Yakuman</h3>
          <p>
            Any hand with 13 or more han from yaku and dora is capped as a
            yakuman
          </p>
        </div>

        <div className="yakuItem">
          <h2>Kokushi musou</h2>
          <hr />
          <h3>Closed: Yakuman</h3>
          <h3>Closed 13 sided wait: Double Yakuman</h3>
          <h3>Open: 0</h3>
          <p>
            Unique hand where the hand has each of the terminal and honor tiles
            plus one more of any terminla or honor tile. If the had is won with
            the 13 sided wait it is a double yakuman 13 Sided wait Kokushi musou
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M1",
              "M9",
              "S1",
              "S9",
              "P1",
              "P9",
              "H1",
              "H2",
              "H3",
              "H4",
              "H5",
              "H6",
              "H7",
              "H7",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Suuankou</h2>
          <hr />
          <h3>Closed: Yakuman</h3>
          <h3>Closed pair wait: Double Yakuman</h3>
          <h3>Open: 0</h3>
          <p>
            Win with 4 closed triplets. If the hand is won when waiting for the
            pair it is a double yakuman Suuankou tanki
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M2",
              "M2",
              "M2",
              "S3",
              "S3",
              "S3",
              "S9",
              "S9",
              "S9",
              "P8",
              "P8",
              "H3",
              "H3",
              "H3",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Chuuren poutou</h2>
          <hr />
          <h3>Closed: Yakuman</h3>
          <h3>Closed 9 sided wait: Double Yakuman</h3>
          <h3>Open: 0</h3>
          <p>
            Win the hand consisting of 1112345678999 plus one of any tile in the
            same suit. If won on the 9 sided wait it becomes the double yakuman
            Junsei chuuren poutou
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M2",
              "M2",
              "M2",
              "S3",
              "S3",
              "S3",
              "S9",
              "S9",
              "S9",
              "P8",
              "P8",
              "H3",
              "H3",
              "H3",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Daisangen</h2>
          <hr />
          <h3>Closed: Yakuman</h3>
          <h3>Open: Yakuman</h3>
          <p>Win with a triplet of each dragon tile</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "H5",
              "H5",
              "H5",
              "H6",
              "H6",
              "H6",
              "H7",
              "H7",
              "H7",
              "B0",
              "B0",
              "B0",
              "B0",
              "B0",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Shousuushii</h2>
          <hr />
          <h3>Closed: Yakuman</h3>
          <h3>Open: Yakuman</h3>
          <p>
            Win with a triplet of 3 of the winds and a pair of the last wind
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "H1",
              "H1",
              "H1",
              "H2",
              "H2",
              "H2",
              "H3",
              "H3",
              "H3",
              "H4",
              "H4",
              "B0",
              "B0",
              "B0",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Daisuushii</h2>
          <hr />
          <h3>Closed: Double Yakuman</h3>
          <h3>Open: Double Yakuman</h3>
          <p>Win with a triplet of each wind tile</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "H1",
              "H1",
              "H1",
              "H2",
              "H2",
              "H2",
              "H3",
              "H3",
              "H3",
              "H4",
              "H4",
              "H4",
              "B0",
              "B0",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Tsuuiisou</h2>
          <hr />
          <h3>Closed: Yakuman</h3>
          <h3>Open: Yakuman</h3>
          <p>Win with only honor tiles</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "H1",
              "H1",
              "H1",
              "H2",
              "H2",
              "H2",
              "H4",
              "H4",
              "H4",
              "H5",
              "H5",
              "H5",
              "H7",
              "H7",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Chinroutou</h2>
          <hr />
          <h3>Closed: Yakuman</h3>
          <h3>Open: Yakuman</h3>
          <p>Win with only terminal tiles</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M1",
              "M1",
              "M1",
              "S1",
              "S1",
              "S1",
              "S9",
              "S9",
              "S9",
              "P1",
              "P1",
              "P1",
              "P9",
              "P9",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Ryuuiisou</h2>
          <hr />
          <h3>Closed: Yakuman</h3>
          <h3>Open: Yakuman</h3>
          <p>
            Win with only tiles that are completely green 2,3,4,6,8 souzu and
            hatsu
          </p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "S2",
              "S2",
              "S2",
              "S3",
              "S4",
              "S6",
              "S6",
              "S6",
              "S8",
              "S8",
              "S8",
              "H6",
              "H6",
              "H6",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Suukantsu</h2>
          <hr />
          <h3>Closed: Yakuman</h3>
          <h3>Open: Yakuman</h3>
          <p>Win with 4 kans and a pair</p>
          <ExampleHand
            className="exampleHand"
            hand={[
              "M2",
              "M2",
              "M2",
              "M2",
              "M8",
              "M8",
              "M8",
              "M8",
              "S7",
              "S7",
              "S7",
              "S7",
              "H3",
              "H3",
              "H3",
              "H3",
              "B0",
              "B0",
            ]}
          />
        </div>

        <div className="yakuItem">
          <h2>Tenhou</h2>
          <hr />
          <h3>Closed: Yakuman</h3>
          <h3>Open: 0</h3>
          <p>Win off the the very first tile drawn as the dealer</p>
        </div>

        <div className="yakuItem">
          <h2>Chiihou</h2>
          <hr />
          <h3>Closed: Yakuman</h3>
          <h3>Open: 0</h3>
          <p>Win off the the very first tile drawn not as dealer</p>
        </div>

        <h2 className="hanTitle">Special</h2>
        <hr />

        <div className="yakuItem">
          <h2>Nagashi mangan</h2>
          <hr />
          <h3>Closed: 5</h3>
          <h3>Open: 5</h3>
          <p>
            If the round ends in a draw and you have only discarded terminal and
            honor tiles and none of your tiles were called, you receive a mangan
          </p>
        </div>
      </div>
    </>
  );
}

function ExampleHand({ hand }) {
  return (
    <>
      <div className="exampleHand">
        {hand.map((tile, index) => (
          <img
            className={"tile"}
            key={index}
            src={`tiles/${tile}.png`}
            draggable="false"
          />
        ))}
      </div>
    </>
  );
}

export default Yaku;
