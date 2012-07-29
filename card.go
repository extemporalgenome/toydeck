// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toydeck

type Card byte

// Returned by the Symbol method when there is no appropriate codepoint to represent the card description.
const Unrepresentable = 0xFFFD

const (
	spades Card = iota + 1
	hearts
	diamonds
	clubs
	nsuits = iota
	black
	white
	ngroups
)

const (
	Spades   = stride * spades
	Hearts   = stride * hearts
	Diamonds = stride * diamonds
	Clubs    = stride * clubs
	Black    = stride * black
	White    = stride * white
	maxord   = stride * ngroups
)

const (
	Unknown Card = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Knight
	Queen
	King
	Joker
	maxrank
)

const (
	maxcard     = maxrank * ngroups
	minsuit     = stride
	maxsuit     = minsuit + maxrank*nsuits
	stride      = 16
	stride52    = 13
	ucs_offset  = 0x1F0A0
	ucs_suit    = 0x2660
	ucs_altsuit = 0x2664
)

var (
	rank2abbr = [maxrank]byte{
		'?',
		'A', '2', '3', '4', '5',
		'6', '7', '8', '9', 'T',
		'J', 'C', 'Q', 'K', '*',
	}

	suit2abbr = [ngroups]byte{
		Unknown: '?', black: 'b', white: 'w',
		spades: 's', hearts: 'h', diamonds: 'd', clubs: 'c',
	}

	rank2name = [maxrank]string{
		"Unranked Card",
		"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "Knight", "King", "Joker",
	}

	suit2name = [ngroups]string{
		Unknown: "Nothing", black: "Black", white: "White",
		spades: "Spades", hearts: "Hearts", diamonds: "Diamonds", clubs: "Clubs",
	}
)

func NewOrd52(ord int) Card {
	if ord < 0 || ord >= 52 {
		return Unknown
	}
	r := Card(ord%stride52 + 1)
	if r >= Knight {
		r++
	}
	return stride + Card(ord/stride52)*stride + r
}

func (c Card) IsValid() bool {
	return c < maxcard
}

func (c Card) IsReal() bool {
	return c > minsuit && (c < maxsuit && c%stride > 0 || c%stride == Joker && c < maxcard)
}

func (c Card) IsPart() bool {
	return c <= minsuit || c < maxcard && (c%stride == 0 || c > maxsuit && c%stride < Joker)
}

func (c Card) Rank() Card {
	if c >= maxcard {
		return Unknown
	}
	return c % stride
}

func (c Card) Suit() Card {
	if s := c / stride; s <= nsuits {
		return s * stride
	}
	return Unknown
}

func (c Card) Color() Card {
	switch c / stride {
	case spades, clubs, black:
		return Black
	case hearts, diamonds, white:
		return White
	}
	return Unknown
}

func (c Card) Ord52() int {
	if c < stride || c >= maxsuit {
		return -1
	}
	r := c % stride
	c -= stride
	switch r {
	case Unknown, Knight, Joker:
		return -1
	case Queen, King:
		r -= 2
	default:
		r -= 1
	}
	return int(c/stride*stride52 + r)
}

func (c Card) GoString() string {
	if c >= maxcard {
		return "XX"
	}
	return string([]byte{rank2abbr[c%stride], suit2abbr[c/stride]})
}

func (c Card) String() string {
	if c >= maxcard {
		return "XX"
	}
	r := string(rank2abbr[c%stride]) // ASCII, so it's okay
	s := c / stride
	switch s {
	case Unknown:
		return r + "?"
	case black:
		return r + "b"
	case white:
		return r + "w"
	}
	return r + string(rune(s-1)+ucs_suit)
}

func (c Card) Name() string {
	if c >= maxcard {
		return "Invalid Card"
	} else if c == Unknown {
		return "Unknown Card"
	}
	s := c / stride
	sn := suit2name[s]
	rn := rank2name[c%stride]
	if s >= black {
		return sn + " " + rn
	}
	return rn + " of " + sn
}

func (c Card) Symbol() rune {
	switch {
	case c == 0:
		return ucs_offset
	case c < stride:
		return rune(rank2abbr[c])
	case c == Black+Joker:
		c = Diamonds + Joker
	case c == White+Joker:
		c = Clubs + Joker
	case c >= maxsuit:
		return Unrepresentable
	case c%stride == 0:
		return rune(c/stride-1) + ucs_suit
	}
	return rune(c-stride) + ucs_offset
}
