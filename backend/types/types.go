package types

type HandObject interface {
}

// point post
type PostHandScore struct {
	Hand         *HandParts        `json:"hand"`
	ScoringParts *HandScoringParts `json:"scoringParts"`
}

type HandParts struct {
	Menzen []string `json:"menzen"`
	Chi    []string `json:"chi"`
	Pon    []string `json:"pon"`
	Kan    []string `json:"kan"`
	Ankan  []string `json:"ankan"`
	Agari  string   `json:"agari"`
}

type HandScoringParts struct {
	Dora    []string `json:"dora"`
	Oya     bool     `json:"oya"`
	Bakaze  string   `json:"bakaze"`
	Jikaze  string   `json:"jikaze"`
	Ron     bool     `json:"ron"`
	Tsumo   bool     `json:"tsumo"`
	Riichi  bool     `json:"riichi"`
	Wriichi bool     `json:"wriichi"`
	Ippatsu bool     `json:"ippatsu"`
	Chankan bool     `json:"chankan"`
	Rinshan bool     `json:"rinshan"`
	Haitei  bool     `json:"haitei"`
	Houtei  bool     `json:"houtei"`
	Tenhou  bool     `json:"tenhou"`
	Chihou  bool     `json:"chihou"`
	Honba   int      `json:"honba"`
	Kiriage bool     `json:"kiriage"`
}

// point return
type ReturnHandScore struct {
	Hand      *HandPartsBlocks `json:"hand"`
	HandScore *Score           `json:"handScore"`
}

type WinningHand struct {
	Hand         [][]string        `json:"hand"`
	Open         bool              `json:"open"`
	ScoringParts *HandScoringParts `json:"scoringParts"`
	HandScore    *Score            `json:"handScore"`
}

type HandPartsBlocks struct {
	Menzen SuitBlocks `json:"menzen"`
	Chi    SuitBlocks `json:"chi"`
	Pon    SuitBlocks `json:"pon"`
	Kan    SuitBlocks `json:"kan"`
	Ankan  SuitBlocks `json:"ankan"`
}

type SuitBlocks struct {
	Manzu [][]string `json:"manzu"`
	Souzu [][]string `json:"souzu"`
	Pinzu [][]string `json:"pinzu"`
	Jihai [][]string `json:"jihai"`
}

type Score struct {
	Han         int                `json:"han"`
	Fu          int                `json:"fu"`
	YakuMan     int                `json:"yakuman"`
	Score       int                `json:"score"`
	ScoreType   string             `json:"scoreType"`
	OyaPayment  int                `json:"oyaPayment"`
	KoPayment   int                `json:"koPayment"`
	YakuList    []*YakuComponet    `json:"yakuList"`
	FuList      []*FuComponet      `json:"fuList"`
	YakumanList []*YakumanComponet `json:"yakumanList"`
}

type YakuComponet struct {
	Title string `json:"title"`
	Han   int    `json:"han"`
}

type FuComponet struct {
	Title string `json:"title"`
	Fu    int    `json:"fu"`
}

type YakumanComponet struct {
	Title   string `json:"title"`
	Yakuman int    `json:"yakuman"`
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
