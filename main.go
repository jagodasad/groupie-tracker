package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

type MyArtistFull struct {
	ID              int                 `json:"id"`
	Image           string              `json:"image"`
	Name            string              `json:"name"`
	Members         []string            `json:"members"`
	CreationDate    int                 `json:"creationDate"`
	FirstAlbum      string              `json:"firstAlbum"`
	Locations       []string            `json:"locations"`
	ConcertDates    []string            `json:"concertDates"`
	DatesLocations  map[string][]string `json:"datesLocations"`
	WikiLink        []string
	TourCity        []string
	TourCountry     []string
	TourDates       [][]string
	TourDateString  []string
	TourDatesString string
}

type MyArtist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type MyLocation struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type MyRelation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type MyDate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type TourData struct {
	ArtistID   int    //artist ID
	RelationID string //key for relations
	City       string
	Country    string
	TourDates  []string
}

type MyTour struct {
	Index []TourData
}

type MyDates struct {
	Index []MyDate `json:"index"`
}

type MyLocations struct {
	Index []MyLocation `json:"index"`
}

type MyRelations struct {
	Index []MyRelation `json:"index"`
}

type MemberWikiLinks struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func main() {
	// static folder
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/concert", concertPage)
	http.HandleFunc("/tour", tourPage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/locations", locationsPage)

	port := ":8080"
	// This allows the port to be changed if needed
	// To change the port use (export PORT=12345) in the command line
	if p, exists := os.LookupEnv("PORT"); exists {
		port = fmt.Sprintf(":%s", p)
	}
	fmt.Printf("Server listening on localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Listen and Serve", err)
	}
}
