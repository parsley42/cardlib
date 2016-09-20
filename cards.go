package cardlib

import (
	"math/rand"
	"time"
)

var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var suits = []string{"c", "d", "h", "s"}
var alltrue = [52]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

type Card int

const NoCard = 52 // if you deal out the whole deck, you get NoCard
const tries = 2

func (c Card) String() string {
	if c == NoCard {
		return ""
	}
	return ranks[c/4] + suits[c%4]
}

func (c Card) GoString() string {
	if c == NoCard {
		return ""
	}
	return ranks[c/4] + suits[c%4]
}

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Deck struct {
	present   [52]bool
	cards     []Card
	full      []Card
	len       int // number of entries in cards
	left      int // cards left in the deck
	refreshes int // how many times cards had to be rebuilt
}

func NewDeck() *Deck {
	full := make([]Card, 52)
	present := alltrue
	cards := full
	for i := 0; i < 52; i++ {
		cards[i] = Card(i)
	}
	return &Deck{
		present: present,
		cards:   cards,
		full:    full,
		len:     52,
		left:    52,
	}
}

func (d *Deck) Reset() {
	d.cards = d.full
	d.present = alltrue
	d.len = 52
	d.left = 52
	d.refreshes = 0
	for i := 0; i < 52; i++ {
		d.cards[i] = Card(i)
	}
}

func (d *Deck) refresh() {
	d.len = d.left
	d.cards = d.full[:d.len]
	i := 0
	for c, p := range d.present {
		if p {
			d.cards[i] = Card(c)
			i++
		}
	}
	d.refreshes++
}

func (d *Deck) Refreshes() int {
	return d.refreshes
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
	for i := 0; i < tries; i++ {
		card := d.cards[r.Intn(d.len)]
		if d.present[card] {
			d.left--
			d.present[card] = false
			return card
		}
	}
	d.refresh()
	card := d.cards[r.Intn(d.len)]
	d.left--
	d.present[card] = false
	return card
}
