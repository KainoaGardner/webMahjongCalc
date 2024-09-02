package hands

type Hand struct {
	Hand string
}

func NewHand() *Hand {
	return &Hand{Hand: "test"}

}
