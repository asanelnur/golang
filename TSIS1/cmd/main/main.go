package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Winners []Winner `json:"winners"`
}

type ResponsePlayer struct {
	Winner Winner `json:"winner"`
}

type Winner struct {
	Year        int    `json:"year"`
	Player      string `json:"player"`
	Nationality string `json:"nationality"`
	Age         int    `json:"age"`
	Club        string `json:"club"`
}

func main() {
	log.Println("starting API server")

	router := mux.NewRouter()
	log.Println("creating routes")
	router.HandleFunc("/health-check", HealthCheck2).Methods("GET")
	router.HandleFunc("/winners", Winners).Methods("GET")
	router.HandleFunc("/winners/{player}", Player).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}

func HealthCheck2(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `Hello, I am developer of this app. I am Elnur Assan.
	This app shows you list of Men's Ballon d'Or award winners using url root: /winners
	And you can see more information about winner using url root: /winners/{player} (For example: winners/Lionel Messi)
	Thanks, Good bye!`)
}

func Winners(w http.ResponseWriter, r *http.Request) {
	log.Println("entering winners end point")
	var response Response
	winners := getWinners()

	response.Winners = winners

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func Player(w http.ResponseWriter, r *http.Request) {
	log.Println("entering player end point")
	vars := mux.Vars(r)
	var response ResponsePlayer
	winner := getWinner(vars["player"])

	if winner.Year == 0 {
		fmt.Fprintf(w, "No such player")
		return
	}

	response.Winner = winner

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func getData() []Winner {
	var winners []Winner

	var winner Winner

	winner.Year = 2015
	winner.Player = "Lionel Messi"
	winner.Nationality = "Argentina"
	winner.Club = "Barcelona"
	winner.Age = 37
	winners = append(winners, winner)

	winner.Year = 2016
	winner.Player = "Cristiano Ronaldo"
	winner.Nationality = "Portugal"
	winner.Club = "Real Madrid"
	winner.Age = 39
	winners = append(winners, winner)

	winner.Year = 2017
	winner.Player = "Cristiano Ronaldo"
	winner.Nationality = "Portugal"
	winner.Club = "Real Madrid"
	winner.Age = 39
	winners = append(winners, winner)

	winner.Year = 2018
	winner.Player = "Luka Modric"
	winner.Nationality = "Croatia"
	winner.Club = "Real Madrid"
	winner.Age = 38
	winners = append(winners, winner)

	winner.Year = 2019
	winner.Player = "Lionel Messi"
	winner.Nationality = "Argentina"
	winner.Club = "Barcelona"
	winner.Age = 37
	winners = append(winners, winner)

	winner.Year = 2021
	winner.Player = "Lionel Messi"
	winner.Nationality = "Argentina"
	winner.Club = "Barcelona"
	winner.Age = 37
	winners = append(winners, winner)
	return winners
}

func getWinners() []Winner {
	return getData()
}

func getWinner(player string) Winner {
	winners := getData()

	for i := 0; i < len(winners); i++ {
		if winners[i].Player == player {
			return winners[i]
		}
	}
	return Winner{}
}
