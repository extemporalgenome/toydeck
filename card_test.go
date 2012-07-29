// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toydeck

import "os"
import "testing"

var tests = []struct {
	card, rank, suit, color Card
	ord52                   int
	gostr, str, name        string
	sym                     rune
}{
	{
		0,
		Unknown,
		Unknown,
		Unknown,
		-1,
		"??",
		"??",
		"Unknown Card",
		'🂠',
	}, {
		Seven,
		7,
		Unknown,
		Unknown,
		-1,
		"7?",
		"7?",
		"Seven of Nothing",
		'7',
	}, {
		Clubs,
		Unknown,
		Clubs,
		Black,
		-1,
		"?c",
		"?♣",
		"Unranked Card of Clubs",
		'♣',
	}, {
		3 + Hearts,
		3,
		Hearts,
		White,
		15,
		"3h",
		"3♡",
		"Three of Hearts",
		'🂳',
	}, {
		King + Diamonds,
		King,
		Diamonds,
		White,
		38,
		"Kd",
		"K♢",
		"King of Diamonds",
		'🃎',
	}, {
		Ace + Spades,
		Ace,
		Spades,
		Black,
		0,
		"As",
		"A♠",
		"Ace of Spades",
		'🂡',
	}, {
		King + Clubs,
		King,
		Clubs,
		Black,
		51,
		"Kc",
		"K♣",
		"King of Clubs",
		'🃞',
	}, {
		Ace + Black,
		Ace,
		Unknown,
		Black,
		-1,
		"Ab",
		"Ab",
		"Black Ace",
		'�',
	}, {
		White + Joker,
		Joker,
		Unknown,
		White,
		-1,
		"*w",
		"*w",
		"White Joker",
		'🃟',
	}, {
		maxcard,
		Unknown,
		Unknown,
		Unknown,
		-1,
		"XX",
		"XX",
		"Invalid Card",
		'�',
	},
}

func TestCards(t *testing.T) {
	for i, test := range tests {
		c := test.card
		gs := c.GoString()
		if v := c.Rank(); v != test.rank {
			t.Error(i, "Expected", gs, "to have the value", test.rank)
		}
		if v := c.Suit(); v != test.suit {
			t.Error(i, "Expected", gs, "to have the suit", test.suit, "but found", v)
		}
		if v := c.Color(); v != test.color {
			t.Error(i, "Expected", gs, "to have the color", test.color, "but found", v)
		}
		if v := c.Ord52(); v != test.ord52 {
			t.Error(i, "Expected", gs, "to have the ordinal value", test.ord52, "but found", v)
		}
		if v := c.GoString(); v != test.gostr {
			t.Error(i, "Expected", gs, "to have the go-string", test.gostr, "but found", v)
		}
		if v := c.String(); v != test.str {
			t.Error(i, "Expected", gs, "to have the string", test.str, "but found", v)
		}
		if v := c.Name(); v != test.name {
			t.Error(i, "Expected", gs, "to have the name", test.name, "but found", v)
		}
		if v := c.Symbol(); v != test.sym {
			t.Error(i, "Expected", gs, "to have the symbol", test.sym, "but found", v)
		}
	}
}

func TestFill(t *testing.T) {
	f, err := os.Create("fill52.txt")
	if err != nil {
		return
	}
	defer f.Close()
	for i := 0; i < 52; i++ {
		c := NewOrd52(i)
		switch i % stride52 {
		case 0:
			f.WriteString(string(c.Suit().Symbol()))
			fallthrough
		default:
			f.WriteString(string(c.Symbol()))
		case stride52 - 1:
			f.WriteString(string(c.Symbol()) + string(c.Suit().Symbol()) + "\n")
		}
	}
}
