package config

<<<<<<< HEAD

=======
>>>>>>> 07fa130c23d628e2c72774158fbefaf535c519e8
import(
	"os"
	mgo "gopkg.in/mgo.v2"
)

<<<<<<< HEAD
func GetMongoDB() (*mgo.Database, error){
	host:= os.Getenv("MONGO_HOST")
	dbName:= os.Getenv("MONGO_DB_NAME")
=======
func getMongoDb() (*mgo.Database, error){
	host:= os.Getenv("MONGO_HOST")
	dbName:= os.Getenv("MONGO_DB_NAMEs")
>>>>>>> 07fa130c23d628e2c72774158fbefaf535c519e8

	session, err := mgo.Dial(host)

	if err != nil{
		return nil, err
	}

	db:= session.DB(dbName)

	return db, nil
}