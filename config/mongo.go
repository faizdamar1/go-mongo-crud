package config

import(
	"os"
	mgo "gopkg.in/mgo.v2"
)

func getMongoDb() (*mgo.Database, error){
	host:= os.Getenv("MONGO_HOST")
	dbName:= os.Getenv("MONGO_DB_NAMEs")

	session, err := mgo.Dial(host)

	if err != nil{
		return nil, err
	}

	db:= session.DB(dbName)

	return db, nil
}