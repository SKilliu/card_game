package models

import "github.com/google/uuid"

type Game struct {
	Players []Player

	Rounds []Round
	// Points of every player(if it is 2x2 game type)
	Points map[string]uint
	// Type can be 2x2 or x1
	Type string
}

type Player struct {
	ID string
}

func NewGame(playersAmount int64, gameType string) *Game {
	game := &Game{
		Type: gameType,
	}

	// create players in a game
	for playersAmount > 0 {
		game.Players = append(game.Players, Player{
			ID: uuid.New().String(),
		})
		playersAmount--
	}

	// init player points
	points := make(map[string]uint)
	for _, player := range game.Players {
		points[player.ID] = 0
	}

	return game
}

func (g *Game) AddRound() {
	deck := GenerateDeck()
	deck.SetTrump(len(g.Players))

	round := Round{
		PlayersCards: nil,
		TurnsOrder:   nil,
		TurnNow:      "",
		Points:       nil,
		Deck:         *deck,
	}
	g.Rounds = append(g.Rounds, round)
}
