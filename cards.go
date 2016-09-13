package cardlib

type Card int

type Deck struct {
	present [52]bool
	left    []Card
}

func NewDeck() *Deck {
	left := make([]Card, 52)
	for i := 0; i < 52; i++ {
		left[i] = Card(i)
	}
	return &Deck{
		present: [52]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
		left:    left,
	}
}

func (d *Deck) hasCard(c Card) bool {
	return d.present[int(c)]
}

func (d *Deck) remove(c Card) {
	d.present[c] = false
	for i := 0; i < len(d.left); i++ {
		if d.left[i] == c {
			d.left = append(d.left[:i], d.left[i+1:]...)
			break
		}
	}
}
