package cardlib

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Card int
type CardRune rune

type Deck struct {
	present   [52]bool
	cards     []Card
	full      []Card
	len       int // number of entries in cards
	left      int // cards left in the deck
	refreshes int // how many times cards had to be rebuilt
}

type CardSpec struct {
	cr      [4]CardRune // Runes specifying a card ("Ah", "2c") or range ("ATs+", "KK", "T9s")
	len     int
	isrange bool
}

var ranks = []CardRune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var suits = []CardRune{'c', 'd', 'h', 's'}

var ranktable map[CardRune]int
var suittable map[CardRune]int

const NoCard = 52 // if you deal out the whole deck, you get NoCard
const tries = 2

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	ranktable = make(map[CardRune]int)
	suittable = make(map[CardRune]int)
	for k, v := range ranks {
		ranktable[v] = k
	}
	for k, v := range suits {
		suittable[v] = k
	}
}

func ParseCard(s string) *CardSpec {
	cs := &CardSpec{}
	l := 0
	for _, r := range s {
		cs.cr[l] = CardRune(r)
		l++
	}
	cs.len = l
	_, cs.isrange = ranktable[cs.cr[1]]
	return cs
}

func (cs *CardSpec) CardNum() Card {
	if cs.isrange {
		return NoCard
	}
	return Card(cs.cr[0].Rank()*4 + cs.cr[1].Suit())
}

func (cs CardRune) Rank() int {
	return ranktable[cs]
}

func (cs CardRune) Suit() int {
	return suittable[cs]
}

func (c Card) RankStr() CardRune {
	return ranks[c/4]
}

func (c Card) SuitStr() CardRune {
	return suits[c%4]
}

func (cs CardRune) String() string {
	return string(cs)
}

func (cs CardRune) GoString() string {
	return string(cs)
}

func (c Card) String() string {
	if c == NoCard {
		return ""
	}
	return (string(ranks[c/4]) + string(suits[c%4]))
}

func (c Card) GoString() string {
	if c == NoCard {
		return ""
	}
	return (string(ranks[c/4]) + string(suits[c%4]))
}

func (d Deck) Print() {
	ranks := make([]string, 52)
	suits := make([]string, 52)
	present := make([]string, 52)
	for i := 0; i < 52; i++ {
		ranks[i] = string(Card(i).RankStr())
		suits[i] = string(Card(i).SuitStr())
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
