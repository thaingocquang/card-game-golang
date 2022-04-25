package dao

import (
	"card-game-golang/model"
	"card-game-golang/module/database"
	"card-game-golang/util"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Player struct{}

// Create ...
func (Player) Create(player model.Player) error {
	var playerCol = database.PlayerCol()

	// InsertOne
	if _, err := playerCol.InsertOne(context.Background(), player); err != nil {
		return err
	}

	return nil
}

// FindByID ...
func (Player) FindByID(ID primitive.ObjectID) (model.Player, error) {
	var (
		playerCol = database.PlayerCol()
		player    model.Player
	)

	// filter
	filter := bson.M{"_id": ID}

	// FindOne
	if err := playerCol.FindOne(context.Background(), filter).Decode(&player); err != nil {
		return player, err
	}

	return player, nil
}

// FindByEmail ...
func (Player) FindByEmail(email string) (model.Player, error) {
	var (
		playerCol = database.PlayerCol()
		player    model.Player
	)

	// filter
	filter := bson.M{"email": email}

	// FindOne
	if err := playerCol.FindOne(context.Background(), filter).Decode(&player); err != nil {
		return model.Player{}, err
	}

	return player, nil
}

// Update ...
func (Player) Update(ID primitive.ObjectID, player model.Player) error {
	var playerCol = database.PlayerCol()

	// UpdateOne
	if _, err := playerCol.UpdateOne(context.Background(), bson.M{"_id": ID}, bson.M{"$set": player}); err != nil {
		return err
	}

	return nil
}

// DeleteByID ...
func (Player) DeleteByID(ID primitive.ObjectID) error {
	var playerCol = database.PlayerCol()

	// filter
	filter := bson.M{"_id": ID}

	dr, err := playerCol.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	fmt.Println(dr.DeletedCount)

	//// DeleteOne ...
	//if dr, err := playerCol.DeleteOne(context.Background(), filter); err != nil {
	//	return err
	//}

	return nil
}

// DeleteAll ...
func (Player) DeleteAll() error {
	var playerCol = database.PlayerCol()

	// DeleteMany
	if _, err := playerCol.DeleteMany(context.Background(), bson.M{}); err != nil {
		return err
	}

	return nil
}

// GetList ...
func (Player) GetList(page, limit int) ([]model.Player, error) {
	var (
		playerCol = database.PlayerCol()
		players   []model.Player
	)

	// options
	opts := new(options.FindOptions)

	if limit != 0 {
		if page == 0 {
			page = 1
		}
		opts.SetSkip(int64((page - 1) * limit))
		opts.SetLimit(int64(limit))
	}

	cursor, err := playerCol.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &players); err != nil {
		return nil, err
	}

	return players, nil
}

// GetListProfile ...
func (Player) GetListProfile(paging *util.Paging) ([]model.Profile, error) {
	var (
		playerCol = database.PlayerCol()
		profiles  []model.Profile
	)

	//// options
	//opts := new(options.AggregateOptions)
	//opts.
	//	SetSkip(int64((paging.Page - 1) * paging.Limit))
	//opts.SetLimit(int64(paging.Limit))

	// count document in playerCol
	count, err := playerCol.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	// set paging total
	paging.Total = count

	// stage
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "stats"}, {"localField", "_id"}, {"foreignField", "playerID"}, {"as", "stat"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$stat"}, {"preserveNullAndEmptyArrays", false}}}}
	skipStage := bson.D{{"$skip", int64((paging.Page - 1) * paging.Limit)}}
	limitStage := bson.D{{"$limit", paging.Limit}}

	// aggregate
	cursor, err := playerCol.Aggregate(context.Background(), mongo.Pipeline{lookupStage, unwindStage, skipStage, limitStage})

	if err = cursor.All(context.Background(), &profiles); err != nil {
		return nil, err
	}

	return profiles, nil
}

func (Player) IsEmailExisted(email string) (bool, error) {
	var playerCol = database.PlayerCol()

	// filter
	filter := bson.M{"email": email}

	//
	count, err := playerCol.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// CountAllPlayer ...
func (Player) CountAllPlayer() int {
	var (
		playerCol = database.PlayerCol()
		ctx       = context.Background()
	)
	count, err := playerCol.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0
	}
	return int(count)
}
