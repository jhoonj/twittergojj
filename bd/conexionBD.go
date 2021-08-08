package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb://twittergo:oxtrNBnABum4WBFcTiPj5gVbu468TxMdR81lScloiOLT5l7yOEaPlfNZmTbhWvrESEoyDLcp9TUoxXJ5XSvs8w==@twittergo.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@twittergo@")

func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions) // context.TODO() significa qeu se cree sin ningun tipo de restriccion
	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client

	}

	log.Println("conexion exitosa con la bd")
	return client
}

func ChequeoConexion() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1

}
