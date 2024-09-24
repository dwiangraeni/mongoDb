package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	db := client.Database("campaign").Collection("voucher")

	//insertData(db)
	//insertMany(db)
	//findOne(db)
	//findAll(db)
	//updateOne(db)
	//updateMany(db)
	//replaceOne(db)
	deleteOne(db)
	//deleteMany(db)
	//findWithFilter(db)
	//findWithProjection(db)
	//findWithSort(db)
	//findWithLimit(db)
	//findWithSkip(db)
	//findWithSkipAndLimit(db)
	//countDocuments(db)
	//distinctValues(db)

}

func insertData(db *mongo.Collection) {
	data := bson.D{
		{"voucher_code", "VOUCHERDWI1"},
		{"min_purchase", 10000},
		{"min_purchase", 5000},
		{"discount_amount", 1000},
		{"max_usage", 100},
	}
	result, err := db.InsertOne(
		context.TODO(),
		data,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Inserted a single document: ", result.InsertedID)
}

func insertMany(db *mongo.Collection) {
	docs := []interface{}{
		bson.D{{"voucher_code", "VOUCHER1"}, {"min_purchase", 1000}},
		bson.D{{"voucher_code", "VOUCHER2"}, {"min_purchase", 2000}, {"discount_amount", 100}},
		bson.D{{"voucher_code", "VOUCHER3"}, {"min_purchase", 3000}, {"discount_amount", 200}, {"max_usage", 100}},
		bson.D{{"voucher_code", "VOUCHER4"}},
	}
	result, err := db.InsertMany(context.TODO(), docs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Inserted multiple documents: ", result.InsertedIDs)
}

func updateOne(db *mongo.Collection) {
	filter := bson.D{{"voucher_code", "VOUCHER1"}}
	update := bson.D{
		{"$set", bson.D{
			{"discount_amount", 200},
		}},
	}
	result, err := db.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}

func updateMany(db *mongo.Collection) {
	filter := bson.D{{"discount_amount", 100}}
	update := bson.D{
		{"$set", bson.D{
			{"discount_amount", 500},
		}},
	}
	result, err := db.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}

func replaceOne(db *mongo.Collection) {
	filter := bson.D{{"voucher_code", "VOUCHER1"}}
	replacement := bson.D{
		{"voucher_code", "VOUCHER1"},
		{"min_purchase", 1500},
		{"discount_amount", 200},
	}
	result, err := db.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Replaced %v Documents!\n", result.ModifiedCount)
}

func deleteOne(db *mongo.Collection) {
	filter := bson.D{{"voucher_code", "VOUCHER1"}}
	result, err := db.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Deleted %v Documents!\n", result.DeletedCount)
}

func deleteMany(db *mongo.Collection) {
	filter := bson.D{{"discount_amount", 500}}
	result, err := db.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Deleted %v Documents!\n", result.DeletedCount)
}

func findOne(db *mongo.Collection) {
	var result bson.M
	data := bson.D{{"voucher_code", "VOUCHERDWI1"}}
	if err := db.FindOne(context.TODO(), data).Decode(&result); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Found a single document: ", result)
}

func findAll(db *mongo.Collection) {
	cursor, err := db.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}

func findWithFilter(db *mongo.Collection) {
	filter := bson.D{{"discount_amount", 200}}
	cursor, err := db.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}

// what is projection?
// Projection is the process of selecting only the necessary data from the database to improve performance.
func findWithProjection(db *mongo.Collection) {
	filter := bson.D{{"discount_amount", 200}}
	projection := bson.D{{"voucher_code", 1}, {"discount_amount", 1}}
	cursor, err := db.Find(context.TODO(), filter, options.Find().SetProjection(projection))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}

func findWithSort(db *mongo.Collection) {
	filter := bson.D{{"discount_amount", 200}}
	sort := bson.D{{"discount_amount", 1}}
	cursor, err := db.Find(context.TODO(), filter, options.Find().SetSort(sort))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}

func findWithLimit(db *mongo.Collection) {
	filter := bson.D{{"discount_amount", 200}}
	limit := int64(1)
	cursor, err := db.Find(context.TODO(), filter, options.Find().SetLimit(limit))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}

func findWithSkip(db *mongo.Collection) {
	filter := bson.D{{"discount_amount", 200}}
	skip := int64(1)
	cursor, err := db.Find(context.TODO(), filter, options.Find().SetSkip(skip))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}

func findWithSkipAndLimit(db *mongo.Collection) {
	filter := bson.D{{"discount_amount", 200}}
	skip := int64(1)
	limit := int64(1)
	cursor, err := db.Find(context.TODO(), filter, options.Find().SetSkip(skip).SetLimit(limit))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}

func countDocuments(db *mongo.Collection) {
	filter := bson.D{{}}
	count, err := db.CountDocuments(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Counted %v Documents!\n", count)
}

func distinctValues(db *mongo.Collection) {
	fieldName := "voucher_code"
	filter := bson.D{{"discount_amount", 200}}
	values, err := db.Distinct(context.TODO(), fieldName, filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Distinct values: ", values)
}
