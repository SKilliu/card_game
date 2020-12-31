package models

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck with cards for round. For every round we need to generate a new deck with a new trump
type Deck struct {
	Cards []Card
	Trump string
}

// SetTrump sets a trump for the deck
func (d *Deck) SetTrump(playersAmount int) {
	trumpCardNumber := playersAmount * HandInitialSize
	d.Trump = d.Cards[trumpCardNumber].Suit

	// we set all cards with trump suit as a trump cards
	for ind := range d.Cards {
		if d.Cards[ind].Suit == d.Trump {
			d.Cards[ind].IsTrump = true
		}
	}

	fmt.Println(d.Cards, d.Trump)
}

// GenerateDeck generates a new deck for a game round
func GenerateDeck() *Deck {
	deck := &Deck{}

	for _, sym := range cardsAscending[:len(cardsAscending)-2] {
		for _, suit := range suits {
			deck.Cards = append(deck.Cards, Card{
				Symbol:  sym,
				Suit:    suit,
				IsTrump: false,
			})
		}
	}

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})

	return deck
}
