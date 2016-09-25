package cardlib

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var suits = []string{"c", "d", "h", "s"}

type Card int

const NoCard = 52 // if you deal out the whole deck, you get NoCard
const tries = 2

func (d Deck) Print() {
	ranks := make([]string, 52)
	suits := make([]string, 52)
	present := make([]string, 52)
	for i := 0; i < 52; i++ {
		ranks[i] = Card(i).Rank()
		suits[i] = Card(i).Suit()
		if d.present[i] {
			present[i] = "t"
		} else {
			present[i] = "f"
		}
	}
	fmt.Println(strings.Join(ranks, " "))
	fmt.Println(strings.Join(suits, " "))
	fmt.Println(strings.Join(present, " "))
	fmt.Println("Cards left:", d.left)
}

func (c Card) Rank() string {
	return ranks[c/4]
}

func (c Card) Suit() string {
	return suits[c%4]
}

func (c Card) String() string {
	if c == NoCard {
		return ""
	}
	return c.Rank() + c.Suit()
}

func (c Card) GoString() string {
	if c == NoCard {
		return ""
	}
	return ranks[c/4] + suits[c%4]
}

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
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
	cards := full
	d := &Deck{
		cards: cards,
		full:  full,
		len:   52,
		left:  52,
	}
	for i := 0; i < 52; i++ {
		cards[i] = Card(i)
		d.present[i] = true
	}
	return d
}

func (d *Deck) Reset() {
	d.cards = d.full
	d.len = 52
	d.left = 52
	d.refreshes = 0
	for i := 0; i < 52; i++ {
		d.cards[i] = Card(i)
		d.present[i] = true
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
		card := d.cards[random.Intn(d.len)]
		if d.present[card] {
			d.left--
			d.present[card] = false
			return card
		}
	}
	d.refresh()
	card := d.cards[random.Intn(d.len)]
	d.left--
	d.present[card] = false
	return card
}
