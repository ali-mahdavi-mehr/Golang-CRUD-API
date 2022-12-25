package mydatabase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	PersonModel "webserverwithGo/Models"
)

type MongoClient struct {
	URL            string
	DBName         string
	CollectionName string
	ctx            context.Context
	client         *mongo.Client
}

func (Mongo *MongoClient) connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	Mongo.ctx = ctx
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(Mongo.URL))
	if err != nil {
		panic(err)
	}
	Mongo.client = client
}

func (Mongo *MongoClient) InsertOne(person PersonModel.Person) (*mongo.InsertOneResult, error) {
	collection := Mongo.client.Database(Mongo.DBName).Collection(Mongo.CollectionName)
	result, err := collection.InsertOne(Mongo.ctx, person)
	return result, err
}

func (Mongo *MongoClient) FindAll() []PersonModel.Person {
	collection := Mongo.client.Database(Mongo.DBName).Collection(Mongo.CollectionName)

	findOptions := options.Find()
	var users []PersonModel.Person

	result, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for result.Next(context.TODO()) {
		var user PersonModel.Person
		err := result.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)

	}
	return users
}

func (Mongo *MongoClient) Close() {
	Mongo.client.Disconnect(Mongo.ctx)
}

func GetMyDb() MongoClient {
	client := MongoClient{
		URL:            "mongodb://localhost:27017",
		DBName:         "MongoTest",
		CollectionName: "Users",
	}
	client.connect()
	return client

}
