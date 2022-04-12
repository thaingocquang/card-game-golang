package model

type (
	// Card ...
	Card struct {
		Name string
		Rank int
		Suit int
	}

	// Hand ...
	Hand struct {
		Cards   []Card
		MaxCard Card
	}
)
