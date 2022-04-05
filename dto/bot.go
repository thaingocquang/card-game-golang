package dto

type (
	// Bot ...
	Bot struct {
		Name         string `json:"name"`
		TotalPoints  int    `json:"totalPoints"`
		RemainPoints int    `json:"remainPoints"`
		MinBet       int    `json:"minBet"`
		MaxBet       int    `json:"maxBet"`
	}
)
