package cardlib

import (
	"math/rand"
	"time"
)

var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var suits = []string{"c", "d", "h", "s"}

type Card int

const NoCard = 52 // if you deal out the whole deck, you get NoCard

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Deck struct {
	present [52]bool
	cards   []Card
	len     int // number of entries in cards
	left    int // cards left in the deck
}

func NewDeck() *Deck {
	cards := make([]Card, 52)
	for i := 0; i < 52; i++ {
		cards[i] = Card(i)
	}
	return &Deck{
		present: [52]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
		cards:   cards,
		len:     52,
		left:    52,
	}
}

func (d *Deck) refresh() {
	d.len = d.left
	d.cards = make([]Card, d.len)
	i := 0
	for c, p := range d.present {
		if p {
			d.cards[i] = Card(c)
			i++
		}
	}
}

func (d *Deck) HasCard(c Card) bool {
	if c == NoCard {
		return false
	}
	return d.present[int(c)]
}

func (d *Deck) Remove(c Card) {
	d.present[c] = false
	d.left--
}

func (d *Deck) Deal() Card {
	if d.left == 0 {
		return NoCard
	}
	card := d.cards[r.Intn(d.len)]
	if d.present[card] {
		d.left--
		d.present[card] = false
		return card
	} else {
		d.refresh()
		card := d.cards[r.Intn(d.len)]
		d.left--
		d.present[card] = false
		return card
	}
}
