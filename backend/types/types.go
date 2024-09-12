package types

type HandObject interface {
}

// point post
type PostHandScore struct {
	Hand    *HandParts `json:"hand"`
	Dora    []string   `json:"dora"`
	Bakaze  string     `json:"bakaze"`
	Jikaze  string     `json:"jikaze"`
	Ron     bool       `json:"ron"`
	Tsumo   bool       `json:"tsumo"`
	Riichi  bool       `json:"riichi"`
	Wriichi bool       `json:"wriichi"`
	Ippatsu bool       `json:"ippatsu"`
	Chankan bool       `json:"chankan"`
	Rinshan bool       `json:"rinshan"`
	Haitei  bool       `json:"haitei"`
	Tenhou  bool       `json:"tenhou"`
	Chihou  bool       `json:"chihou"`
	Renchan int        `json:"renchan"`
}

type HandParts struct {
	Menzen []string `json:"menzen"`
	Chi    []string `json:"chi"`
	Pon    []string `json:"pon"`
	Kan    []string `json:"kan"`
	Ankan  []string `json:"ankan"`
}

type HandPartsBlocks struct {
	Menzen Suits `json:"menzen"`
	Chi    Suits `json:"chi"`
	Pon    Suits `json:"pon"`
	Kan    Suits `json:"kan"`
	Ankan  Suits `json:"ankan"`
}

type Suits struct {
	Manzu [][]string `json:"manzu"`
	Souzu [][]string `json:"souzu"`
	Pinzu [][]string `json:"pinzu"`
	Jihai [][]string `json:"jihai"`
}

// point return
type ReturnHandScore struct {
	Hand      *HandPartsBlocks `json:"hand"`
	HandScore *Score           `json:"handscore"`
}

type Score struct {
	Han        int             `json:"han"`
	Fu         int             `json:"fu"`
	Score      int             `json:"score"`
	ScoreType  string          `json:"scoretype"`
	Yaku       []*YakuComponet `json:"yaku"`
	FuComponet []*FuComponet   `json:"fucomponet"`
}

type YakuComponet struct {
	Title string `json:"title"`
	Han   int    `json:"han"`
}

type FuComponet struct {
	Title  string `json:"title"`
	Amount int    `json:"Amount"`
}

// H (1,2,3,4 = ton,nan,sha,pei)  (5 6 7 = haku,hatsu,chun)
// M manzu  S souzu P pinzu
// A aka
var Tiles = map[string]bool{
	"H1": true,
	"H2": true,
	"H3": true,
	"H4": true,
	"H5": true,
	"H6": true,
	"H7": true,

	"M1":  true,
	"M2":  true,
	"M3":  true,
	"M4":  true,
	"M5":  true,
	"M6":  true,
	"M7":  true,
	"M8":  true,
	"M9":  true,
	"M5A": true,

	"S1":  true,
	"S2":  true,
	"S3":  true,
	"S4":  true,
	"S5":  true,
	"S6":  true,
	"S7":  true,
	"S8":  true,
	"S9":  true,
	"S5A": true,

	"P1":  true,
	"P2":  true,
	"P3":  true,
	"P4":  true,
	"P5":  true,
	"P6":  true,
	"P7":  true,
	"P8":  true,
	"P9":  true,
	"P5A": true,
}
