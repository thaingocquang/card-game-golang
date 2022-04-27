package dto

import (
	"card-game-golang/model"
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"strconv"
	"time"
)

const (
	TotalSuit = 4
)

type (
	// GameJSON ...
	GameJSON struct {
		ID         primitive.ObjectID `json:"id"`
		GameNo     int                `json:"gameNo"`
		PlayerID   primitive.ObjectID `json:"playerID"`
		BotID      primitive.ObjectID `json:"botID"`
		WinnerID   primitive.ObjectID `json:"winnerID"`
		WinnerName string             `json:"winnerName"`
		PlayerHand Hand               `json:"playerHand"`
		BotHand    Hand               `json:"botHand"`
		BetValue   int                `json:"betValue"`
		CreatedAt  time.Time          `json:"createdAt"`
		UpdatedAt  time.Time          `json:"updatedAt"`
	}

	// GameVal ...
	GameVal struct {
		BetValue int `json:"betValue"`
	}

	// Card ...
	Card struct {
		Name string `json:"name,omitempty"`
		Rank int    `json:"rank,omitempty"`
		Suit int    `json:"suit,omitempty"`
	}

	// Hand ...
	Hand struct {
		Cards   []Card `json:"cards"`
		MaxCard Card   `bson:"maxCard"`
	}

	// DeckCard ...
	DeckCard []Card
)

// ConvertToBSON ...
func (h Hand) ConvertToBSON() model.Hand {
	var handBSON model.Hand
	for _, v := range h.Cards {
		handBSON.Cards = append(handBSON.Cards, model.Card{
			Name: v.Name,
			Rank: v.Rank,
			Suit: v.Suit,
		})
	}
	handBSON.MaxCard = model.Card{
		Name: h.MaxCard.Name,
		Rank: h.MaxCard.Rank,
		Suit: h.MaxCard.Suit,
	}

	return handBSON
}

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

	baTien := true
	for _, c := range h.Cards {
		if c.Rank < 11 || c.Rank == 14 {
			baTien = false
		}
	}

	if baTien == true {
		return 999
	}

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
