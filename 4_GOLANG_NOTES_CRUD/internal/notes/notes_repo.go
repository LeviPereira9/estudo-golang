package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// repo -> data access layer
type Repo struct {
	coll *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo{
	return &Repo{
		coll: db.Collection("notes"),
	}
}

func (r *Repo) Create(ctx context.Context, note Note) (Note, error){
	opCtx, cancel := context.WithTimeout(ctx, 5 *time.Second)

	defer cancel()

	_, err := r.coll.InsertOne(opCtx, note)

	if err != nil{
		return Note{}, fmt.Errorf("Insert note failed")
	}


	return note, nil
}

func (r *Repo) List(ctx context.Context) ([]Note, error){

	opCtx, cancel := context.WithTimeout(ctx, 5 *time.Second)
	defer cancel()

	filter := bson.M{}//match all docs

	// Fidn return a cursor (like an iterator) -> over matching elements
	cursors, err := r.coll.Find(opCtx, filter)

	if err != nil{
		return nil, fmt.Errorf("find notes failed")
	}

	// cursor must be closed after use
	// avoid any kind of leaks 
	defer cursors.Close(opCtx)
	
	var notes []Note

	if err := cursors.All(opCtx, &notes); err != nil {
		return nil, fmt.Errorf("Decode noes failed")
	}

	return notes, nil
}

func (r *Repo) GetByiD(ctx context.Context, id bson.ObjectID)(Note, error){
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	var note Note

	err := r.coll.FindOne(opCtx, filter, options.FindOne()).Decode(&note)

	if err != nil {
		return Note{}, fmt.Errorf("Find note by id failed")
	}

	return note, nil
}

func (r *Repo) UpdateById(ctx context.Context, id bson.ObjectID, req UpdateNoteRequest) (Note, error){

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"title": req.Title,
			"content": req.Content,
			"pinned": req.Pinned,
			"updatedAt": time.Now().UTC(),
		},
	}

	after := options.After

	opts := options.FindOneAndUpdate().SetReturnDocument(after)

	var updated Note
	
	err := r.coll.FindOneAndUpdate(opCtx, filter, update, opts).Decode(&updated)

	if err != nil{
		return Note{}, fmt.Errorf("updated note failed")
	}
	
	return updated, nil
}

func (r *Repo) DeleteByID(ctx context.Context, id bson.ObjectID) (bool, error){
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	res, err := r.coll.DeleteOne(opCtx, filter)

	if err != nil {
		return false, fmt.Errorf("Failed to delete the given note")
	}

	if res.DeletedCount == 0 {
		return false, nil
	}

	return true, nil
}

func (r *Repo) DeleteNoteByID(ctx context.Context, id bson.ObjectID) (bool, error){

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()
	
	filter := bson.M{"_id": id}

	res, err := r.coll.DeleteOne(opCtx, filter)

	if err != nil{
		return false, fmt.Errorf("failed to delete the given note")
	}

	if res.DeletedCount == 0{
		return false, nil	
	}

	return true, nil
	
}