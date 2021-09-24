package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"writerxl-api/data"
)

type Profile struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Email       string             `json:"email"`
	Nickname    string             `bson:"nickname,omitempty" json:"nickname"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Picture     string             `bson:"picture,omitempty" json:"picture"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
}

func CreateProfile(profile Profile) error {
	client, err := data.GetMongoClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), data.DefaultTimeout)
	defer cancel()

	collection := client.Database(data.DB).Collection(data.PROFILE)

	profile.ID = primitive.NewObjectID()
	_, err = collection.InsertOne(ctx, profile)
	if err != nil {
		return err
	}

	return nil
}

func GetProfileById(id string) (Profile, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != err {
		return Profile{}, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}
	return getProfile(filter)
}

func GetProfileByEmail(email string) (Profile, error) {
	filter := bson.D{primitive.E{Key: "email", Value: email}}
	return getProfile(filter)
}

func getProfile(filter bson.D) (Profile, error) {
	result := Profile{}

	client, err := data.GetMongoClient()
	if err != nil {
		return result, err
	}

	collection := client.Database(data.DB).Collection(data.PROFILE)

	ctx, cancel := context.WithTimeout(context.Background(), data.DefaultTimeout)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func UpdateProfile(id string, profile Profile) (Profile, error) {
	doc := Profile{}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != err {
		return doc, err
	}

	client, err := data.GetMongoClient()
	if err != nil {
		return doc, err
	}

	collection := client.Database(data.DB).Collection(data.PROFILE)

	ctx, cancel := context.WithTimeout(context.Background(), data.DefaultTimeout)
	defer cancel()

	filter := bson.M{"_id": objectId}

	update := bson.M{
		"$set": bson.M{
			"email":       profile.Email,
			"nickname":    profile.Nickname,
			"name":        profile.Name,
			"picture":     profile.Picture,
			"description": profile.Description,
		},
	}

	after := options.After
	upsert := false

	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := collection.FindOneAndUpdate(ctx, filter, update, &opts)
	if result.Err() != nil {
		return Profile{}, result.Err()
	}
	err = result.Decode(&doc)

	return doc, err
}
