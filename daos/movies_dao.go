package daos

import (
	. "github.com/mlabouardy/movies-restapi/models"
	"gopkg.in/mgo.v2"
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

func (m *MoviesDAO) connect(){

	session, error := mgo.Dial(m.Server)
	if error != nil{
		log.Fatal(error)
	}

	db := session.DB(m.Database)
}


func (m *MoviesDAO)insert(movie Movie) error{
	e := db.C(COLLECTION).Insert(&movie)
	return e
}

