package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ddld93/abedmis/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB_Connect struct {
	Session *mongo.Client
}

var (
	database   = "abedmis"
	collection = "beneficiaries"
)

func ConnectDB(host string, port int) *DB_Connect {
	URI := fmt.Sprintf("mongodb://%s:%v", host, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return &DB_Connect{Session: client}
}

func (s *DB_Connect) AddOne(beneficiary models.Beneficiary) (*mongo.InsertOneResult, error) {
	beneficiary.Id = primitive.NewObjectID()
	beneficiary.CreatedAt = time.Now()
	result, err := s.Session.Database(database).Collection(collection).InsertOne(context.TODO(), beneficiary)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *DB_Connect) AddMany(beneficiary []models.Beneficiary) (*mongo.InsertManyResult, error) {
	var beneficiaries []interface{}
	for _, v := range beneficiary {
		v.CreatedAt = time.Now()
		v.Id = primitive.NewObjectID()
		beneficiaries = append(beneficiaries, bson.M{"name":"Alice"},)
	}
	result, err := s.Session.Database(database).Collection(collection).InsertMany(context.TODO(), beneficiaries)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *DB_Connect) GetAll(pspKey string, mode string) ([]models.Beneficiary, error) {
	var beneficiaries []models.Beneficiary
	result, err := s.Session.Database(database).Collection(collection).Find(context.TODO(), bson.M{"psp":pspKey,"mode":mode})
	if err != nil {
		return nil, err
	}
	err = result.All(context.TODO(), &beneficiaries)
	if err != nil {
		return nil, err
	}
	return beneficiaries, nil
}
func (s *DB_Connect) GetOne(id primitive.ObjectID) (models.Beneficiary,error){
	var beneficiaries models.Beneficiary
	err:= s.Session.Database(database).Collection(collection).FindOne(context.TODO(), bson.M{"_id":id}).Decode(&beneficiaries)
	if err != nil {
		return beneficiaries, err
	}
	return beneficiaries,nil
}
