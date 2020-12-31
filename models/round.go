package models

type Round struct {
	PlayersCards map[string][]Card
	TurnsOrder   []Player
	TurnNow      string
	Points       map[string]uint
	Deck         Deck
}
