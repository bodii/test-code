package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongoDB() (client *mongo.Client) {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到Mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查链接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return
}

func openDatabase(client *mongo.Client, dbname string) (db *mongo.Database) {
	db = client.Database(dbname)
	return
}

func openTable(db *mongo.Database, tablename string) (collection *mongo.Collection) {
	collection = db.Collection(tablename)
	return
}

func close(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func insertOne(collection *mongo.Collection, student Student) {
	insertResult, err := collection.InsertOne(context.TODO(), student)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document:", insertResult.InsertedID)
}

func insertMary(collection *mongo.Collection, students []interface{}) {
	insertResult, err := collection.InsertMany(context.TODO(), students)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents:", insertResult.InsertedIDs)
}

func updateOne(collection *mongo.Collection, where bson.D, update bson.D) {
	updateResult, err := collection.UpdateOne(context.TODO(), where, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"matched %v documents and updated %v documents.\n",
		updateResult.MatchedCount, updateResult.ModifiedCount)
}

func findOne(collection *mongo.Collection, where bson.D) {
	var result Student
	err := collection.FindOne(context.TODO(), where).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document:%+v\n", result)
}

func find(collection *mongo.Collection, limit int64, where bson.D) {
	// find multiple
	findOptions := options.Find()
	findOptions.SetLimit(limit)

	// init result slice
	var results []*Student

	cur, err := collection.Find(context.TODO(), where, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// loop
	for cur.Next(context.TODO()) {
		var elem Student
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// close
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %#v\n", results)
}

func deleteOne(collection *mongo.Collection, where bson.D) {
	result, err := collection.DeleteOne(context.TODO(), where)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v document in the trainers collection\n", result.DeletedCount)
}

func deleteMany(collection *mongo.Collection, where bson.D) {
	result, err := collection.DeleteMany(context.TODO(), where)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", result.DeletedCount)
}

// Student type
type Student struct {
	Name string
	Age  int
}

func main() {
	conn := connectMongoDB()
	defer close(conn)
	db := openDatabase(conn, "test")
	student := openTable(db, "student")

	// s1 := Student{"jack", 18}
	// s2 := Student{"tom", 20}
	// s3 := Student{"lili", 19}

	// insertOne(student, s1)
	// insertMary(student, []interface{}{s2, s3})

	// where := bson.D{{"name", "jack"}}
	// update := bson.D{
	// 	{"$inc", bson.D{
	// 		{"age", 1},
	// 	}},
	// }
	// updateOne(student, where, update)

	// where := bson.D{{"name", "jack"}}
	// findOne(student, where)
	// deleteOne(student, where)

	// delete all
	// where := bson.D{}
	// deleteMany(student, where)

	find(student, 2, bson.D{})

}
