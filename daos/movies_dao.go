package daos

import (
	"gopkg.in/mgo.v2"
	"MoviesApp/models"
	"log"
)

type MoviesDAO struct{
	Server string
	Database string
}

var db *mgo.Database


const(
	COLLECTION = "movies"
)

func (m *MoviesDAO) Connect(){

	session, error := mgo.Dial(m.Server)
	if error != nil{
		log.Fatal(error)
	}

	db = session.DB(m.Database)
}


// Insert a movie into database
func (m *MoviesDAO) Insert(movie models.Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}