package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoC variable que contiene la conexion a la bd
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://jccovarrubias:C3sarcova3@cluster0.ukywijp.mongodb.net/?retryWrites=true&w=majority")

// ConectarBD es la funcion que me permite conectar la bd
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error al hacer ping", err.Error())
		return client
	}
	log.Println("Conexcion establecida a la BD")
	return client
}

// ChequeoConnection verifica la conexion a la bd
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
