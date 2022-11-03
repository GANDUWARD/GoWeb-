package main

import (
	"gitee.com/shirdon1/goWebActualCombat/chapter4/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var (
		client     = mongodb.MgoCli()
		db         *mongo.Database
		collection *mongo.Collection
	)
	//选择数据库my_db
	db = client.Database("my_db")

	//选择表my_collection
	collection = db.Collection("my_collection")
	collection = collection
}
