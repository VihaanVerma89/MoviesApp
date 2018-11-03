package main

import (
	. "MoviesApp/config"
	"MoviesApp/daos"
	"MoviesApp/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

var config = Config{}
var dao  = daos.MoviesDAO{}

func init(){
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
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

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	movie, err := dao.FindById(vars["id"])
	if err!= nil{
		respondWithError(w, http.StatusNotFound, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}


func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie
	if err:=json.NewDecoder(r.Body).Decode(&movie); err != nil{
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err:= dao.Delete(movie) ; err !=nil{
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result":"success"})
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie
	if err:= json.NewDecoder(r.Body).Decode(&movie); err!=nil{
		respondWithError(w, http.StatusBadRequest, "Invalid r payload")
		return
	}

	if update := dao.Update(movie) ; update!=nil{
		respondWithJson(w, http.StatusInternalServerError, update.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result":"success"})
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie
	if error := json.NewDecoder(r.Body).Decode(&movie); error != nil{
		respondWithError(w, http.StatusBadRequest, "Invalid r payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dao.Insert(movie); err != nil{
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}){
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	movies, e := dao.FindAll()

	if e!=nil{
		respondWithError(w, http.StatusNotFound, e.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}





































































































