package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"math/rand"
	"strconv"
)

const (
	TotalSuit = 4
)

type (
	// GameVal ...
	GameVal struct {
		BetValue int `json:"betValue"`
	}

	// Card ...
	Card struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
		Rank int    `json:"rank,omitempty" bson:"rank,omitempty"`
		Suit int    `json:"suit,omitempty" bson:"suit,omitempty"`
	}

	// Hand ...
	Hand struct {
		Cards   []Card
		MaxCard Card
	}

	// DeckCard ...
	DeckCard []Card
)

// Validate ...
func (g GameVal) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.BetValue, validation.Required, validation.Min(0)),
	)
}

// CompareSuit ...
func (c Card) CompareSuit(c1 Card) bool {
	return c.Suit > c1.Suit
}

// SumRank ...
func (h Hand) SumRank() int {
	sumRank := 0
	for _, c := range h.Cards {
		if c.Rank == 14 {
			sumRank += 1
		} else {
			if c.Rank > 10 {
				sumRank += 10
			} else {
				sumRank += c.Rank
			}
		}
	}
	return sumRank % 10
}

// InitMaxCard ...
func (h *Hand) InitMaxCard() {
	h.MaxCard = h.Cards[0]
	if h.Cards[1].Rank > h.Cards[0].Rank {
		h.MaxCard = h.Cards[1]
	}
	if h.Cards[2].Rank > h.Cards[1].Rank {
		h.MaxCard = h.Cards[2]
	}
}

// CompareHandIsHigher ...
func (h Hand) CompareHandIsHigher(h1 Hand) bool {
	hSumRank := h.SumRank()
	h1SumRank := h1.SumRank()

	if hSumRank == h1SumRank {
		if h.MaxCard.Rank == h1.MaxCard.Rank {
			return h.MaxCard.CompareSuit(h1.MaxCard)
		}
		return h.MaxCard.Rank > h1.MaxCard.Rank
	}

	return hSumRank > h1SumRank
}

// Init ...
func (d *DeckCard) Init() {
	// rank: 11 = J, 12 = Q, 13 = K, 14 = A
	// suit: 4 = HEARTS, 3 = DIAMOND, 2 = CLUBES, 1 = SPADES
	for i := 1; i <= TotalSuit; i++ {
		for j := 2; j <= 10; j++ {
			card := Card{
				Name: strconv.Itoa(j),
				Rank: j,
				Suit: i,
			}
			*d = append(*d, card)
		}
		jCard := Card{
			Name: "J",
			Rank: 11,
			Suit: i,
		}
		*d = append(*d, jCard)
		qCard := Card{
			Name: "Q",
			Rank: 12,
			Suit: i,
		}
		*d = append(*d, qCard)
		kCard := Card{
			Name: "K",
			Rank: 13,
			Suit: i,
		}
		*d = append(*d, kCard)
		aCard := Card{
			Name: "A",
			Rank: 14,
			Suit: i,
		}
		*d = append(*d, aCard)
	}
}

// Shuffle ...
func (d *DeckCard) Shuffle() {
	for i := 1; i < len(*d); i++ {
		r := rand.Intn(i + 1)

		if i != r {
			(*d)[r], (*d)[i] = (*d)[i], (*d)[r]
		}
	}
}
