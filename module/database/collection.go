package database

import "go.mongodb.org/mongo-driver/mongo"

const (
	players = "players"
	admin   = "admin"
	bots    = "bots"
	games   = "games"
	stats   = "stats"
)

// PlayerCol ...
func PlayerCol() *mongo.Collection {
	return db.Collection(players)
}

// AdminCol ...
func AdminCol() *mongo.Collection {
	return db.Collection(admin)
}

// BotCol ...
func BotCol() *mongo.Collection {
	return db.Collection(bots)
}

// GameCol ...
func GameCol() *mongo.Collection {
	return db.Collection(games)
}

// StatsCol ...
func StatsCol() *mongo.Collection {
	return db.Collection(stats)
}
