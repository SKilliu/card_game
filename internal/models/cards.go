package models

// Card represent a card in deck
type Card struct {
	Symbol  string
	Suit    string
	IsTrump bool
}

// CardValue represents a card value as a symbol and suit
type CardValue struct {
	Symbol string
	Suit   string
}
