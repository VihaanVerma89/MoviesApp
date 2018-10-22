package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mlabouardy/movies-restapi/models"
	. "github.com/mlabouardy/movies-restapi/dao"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

var dao  = MoviesDAO{}

func init(){
	fmt.Println("init called")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")

	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}

func FindMovieEndpoint(writer http.ResponseWriter, request *http.Request) {

}
func DeleteMovieEndPoint(writer http.ResponseWriter, request *http.Request) {

}
func UpdateMovieEndPoint(writer http.ResponseWriter, request *http.Request) {

}
func CreateMovieEndPoint(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var movie models.Movie
	if error := json.NewDecoder(request.Body).Decode(&movie); error != nil{
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dao.Insert(movie); err != nil{
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(writer, http.StatusCreated, movie)
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(writer http.ResponseWriter, code int, payload interface{}){
	response, _ := json.Marshal(payload)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(response)
}

func AllMoviesEndPoint(writer http.ResponseWriter, request *http.Request) {

}





































































































