package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/mail"
	"time"
	"writerxl-api/data"
)

type Profile struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Email       string             `json:"email"`
	Nickname    string             `json:"nickname,omitempty"`
	Name        string             `json:"name,omitempty"`
	Picture     string             `json:"picture,omitempty"`
	Description string             `json:"description,omitempty"`
	Active      bool               `json:"active"`
	CreatedDate time.Time          `json:"createdDate"`
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
	profile.Active = true
	profile.CreatedDate = time.Now()

	_, err = collection.InsertOne(ctx, profile)
	if err != nil {
		return Profile{}, err
	}

	return GetProfileByEmail(profile.Email)
}

func GetProfileByEmail(email string) (Profile, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return Profile{}, err
	}

	filter := bson.D{primitive.E{Key: "email", Value: email}}
	return getProfile(filter)
}

func GetProfileById(id string) (Profile, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Profile{}, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}
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

func UpsertProfile(id string, profile Profile) (Profile, error) {
	update := bson.M{
		"$set": bson.M{
			"nickname":    profile.Nickname,
			"name":        profile.Name,
			"picture":     profile.Picture,
			"description": profile.Description,
		},
	}

	return updateProfile(id, update)
}

func ActivateProfile(id string) error {
	update := bson.M{
		"$set": bson.M{
			"active": true,
		},
	}

	_, err := updateProfile(id, update)
	return err
}

func DeactivateProfile(id string) error {
	update := bson.M{
		"$set": bson.M{
			"active": false,
		},
	}

	_, err := updateProfile(id, update)
	return err
}

func updateProfile(id string, update bson.M) (Profile, error) {
	doc := Profile{}

	client, err := data.GetMongoClient()
	if err != nil {
		return doc, err
	}

	collection := client.Database(data.DB).Collection(data.PROFILE)

	ctx, cancel := context.WithTimeout(context.Background(), data.DefaultTimeout)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Profile{}, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

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
