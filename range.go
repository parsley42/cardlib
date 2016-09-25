package cardlib

type CardSpec struct {
	cr      [4]CardRune // Runes specifying a card ("Ah", "2c") or range ("ATs+", "KK", "T9s")
	len     int
	isrange bool
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

var pairs = [6][2]CardRune{{'c', 'd'}, {'c', 'h'}, {'c', 's'}, {'d', 'h'}, {'d', 's'}, {'s', 'h'}}

func (c *CardSpec) Range() [][2]Card {
	var r [][2]Card
	if !c.isrange {
		return r
	}
	r1 := c.cr[0].Rank()
	r2 := c.cr[1].Rank()
	switch {
	case r1 == r2: // pair
		r = make([][2]Card, 6)
	}
	return r
}
