package types

type HandObject interface {
}

type PostHandScore struct {
	HandParts *HandParts `json:"handParts"`
}

type HandParts struct {
	Menzen []string `json:"menzen"`
	Chi    []string `json:"chi"`
	Pon    []string `json:"pon"`
	Kan    []string `json:"kan"`
	Ankan  []string `json:"ankan"`
}
