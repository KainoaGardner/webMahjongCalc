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
	Menzen [][]string `json:"menzen"`
	Chi    [][]string `json:"chi"`
	Pon    [][]string `json:"pon"`
	Kan    [][]string `json:"kan"`
	Ankan  [][]string `json:"ankan"`
}

// point return
type ReturnHandScore struct {
	Hand       *HandPartsBlocks `json:"hand"`
	Han        int              `json:"han"`
	Fu         int              `json:"fu"`
	Yaku       []*YakuComponet  `json:"yaku"`
	FuComponet []*FuComponet    `json:"fucomponet"`
}

type YakuComponet struct {
	Title string `json:"title"`
	Han   int    `json:"han"`
}

type FuComponet struct {
	Title  string `json:"title"`
	Amount int    `json:"Amount"`
}
