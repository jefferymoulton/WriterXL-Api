package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"writerxl-api/data"
)

type Profile struct {
	ID          primitive.ObjectID `bson:"_id" json:"-"`
	Email       string             `json:"email,omitempty"`
	Nickname    string             `json:"nickname,omitempty"`
	Name        string             `json:"name,omitempty"`
	Picture     string             `json:"picture,omitempty"`
	Description string             `json:"description,omitempty"`
}

func CreateProfile(profile Profile) (Profile, error) {
	client, err := data.GetMongoClient()
	if err != nil {
		return Profile{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), data.DefaultTimeout)
	defer cancel()

	collection := client.Database(data.DB).Collection(data.PROFILE)

	profile.ID = primitive.NewObjectID()
	_, err = collection.InsertOne(ctx, profile)
	if err != nil {
		return Profile{}, err
	}

	return GetProfileByEmail(profile.Email)
}

func GetProfileByEmail(email string) (Profile, error) {
	filter := bson.D{primitive.E{Key: "email", Value: email}}

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

func UpdateProfile(email string, profile Profile) (Profile, error) {
	doc := Profile{}

	client, err := data.GetMongoClient()
	if err != nil {
		return doc, err
	}

	collection := client.Database(data.DB).Collection(data.PROFILE)

	ctx, cancel := context.WithTimeout(context.Background(), data.DefaultTimeout)
	defer cancel()

	filter := bson.M{"email": email}

	update := bson.M{
		"$set": bson.M{
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
