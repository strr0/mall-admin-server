package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mall-admin-server/common/model"
	"time"
)

type MemberReadHistoryService struct {
	Collection *mongo.Collection
}

func (iService MemberReadHistoryService) Create(history model.MemberReadHistory) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	history.CreateDate = time.Now()
	_, err := iService.Collection.InsertOne(ctx, &history)
	return err
}

func (iService MemberReadHistoryService) Delete(ids []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := iService.Collection.DeleteMany(ctx, bson.M{"id": bson.M{"$in": ids}})
	return err
}

func (iService MemberReadHistoryService) List(memberId string) []model.MemberReadHistory {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	find, err := iService.Collection.Find(ctx, bson.M{"memberid": memberId})
	if err != nil {
		return nil
	}
	results := make([]model.MemberReadHistory, 0)
	for find.Next(ctx) {
		var result model.MemberReadHistory
		err := find.Decode(&result)
		if err == nil {
			results = append(results, result)
		}
	}
	_ = find.Close(ctx)
	return results
}
