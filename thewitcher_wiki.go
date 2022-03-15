package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Wiki struct {
	Data []struct {
		Characters      string `json:"characters"`
		Creatures       string `json:"name"`
		Races           string `json:"races"`
		Species         string `json:"species"`
		Suseptibilities string `json:"suseptibilities"`
	} `json:"results"`
}
type Characters struct {
	Id     int       `json:"id"`
	Name   string    `json:"name"`
	Race   string    `json:"race"`
	Gender string    `json:"gender"`
	Edited time.Time `json:"edited"`
	Url    string    `json:"url"`
}
type Creatures struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Species          string    `json:"species"`
	Susceptibilities []string  `json:"susceptibilities"`
	Edited           time.Time `json:"edited"`
	Url              string    `json:"url"`
}
type Races struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Edited          time.Time `json:"edited"`
	Representatives []string  `json:"representatives"`
}
type Species struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Edited          time.Time `json:"edited"`
	Representatives []string  `json:"representatives"`
}
type Susceptibilities struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Edited          time.Time `json:"edited"`
	Representatives []string  `json:"representatives"`
}

func main() {

	fmt.Println("server starting")

	cssFolder := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", cssFolder))

	imageFolder := http.FileServer(http.Dir("Image"))
	http.Handle("/img/", http.StripPrefix("/img/", imageFolder))

	songFolder := http.FileServer(http.Dir("Audio"))
	http.Handle("/song/", http.StripPrefix("/song/", songFolder))

	//Fonction Handle page principale
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := "https://www.witcher.red/api/"

		httpClient := http.Client{
			Timeout: time.Second * 5, // define timeout
		}

		//create request
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "seb go tuto v2")

		//make api call
		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		//parse response
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		response := Wiki{}
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		//print

		tmpl := template.Must(template.ParseFiles("../The_Witcher/html/index.html"))
		fmt.Println("call server !!!!")
		tmpl.Execute(w, response)
	})

	//Fonction Handle Characters
	http.HandleFunc("/perso", func(w http.ResponseWriter, r *http.Request) {

		var characters []Characters
		url := "https://www.witcher.red/api/characters/"

		httpClient := http.Client{
			//Timeout: time.Second * 2, // define timeout
		}

		//create request
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "seb go tuto v2")

		//make api call
		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		//parse response
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		response := characters
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		//print

		tmpl := template.Must(template.ParseFiles("../The_Witcher/html/perso.html"))
		fmt.Println("call server !!!!")
		tmpl.Execute(w, response)
	})

	//Fonction Handle Creatures
	http.HandleFunc("/creatures", func(w http.ResponseWriter, r *http.Request) {

		var creatures []Creatures
		url := "https://www.witcher.red/api/creatures/"

		httpClient := http.Client{
			//Timeout: time.Second * 2, // define timeout
		}

		//create request
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "seb go tuto v2")

		//make api call
		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		//parse response
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		response := creatures
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		//print

		tmpl := template.Must(template.ParseFiles("../The_Witcher/html/creatures.html"))
		fmt.Println("call server !!!!")
		tmpl.Execute(w, response)
	})

	//Fonction Handle Races
	http.HandleFunc("/races", func(w http.ResponseWriter, r *http.Request) {

		var races []Races
		url := "https://www.witcher.red/api/races/"

		httpClient := http.Client{
			//Timeout: time.Second * 2, // define timeout
		}

		//create request
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "seb go tuto v2")

		//make api call
		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		//parse response
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		response := races
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		//print

		tmpl := template.Must(template.ParseFiles("../The_Witcher/html/races.html"))
		fmt.Println("call server !!!!")
		tmpl.Execute(w, response)
	})

	//Fonction Handle Species
	http.HandleFunc("/species", func(w http.ResponseWriter, r *http.Request) {

		var species []Species
		url := "https://www.witcher.red/api/species/"

		httpClient := http.Client{
			//Timeout: time.Second * 2, // define timeout
		}

		//create request
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "seb go tuto v2")

		//make api call
		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		//parse response
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		response := species
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		//print

		tmpl := template.Must(template.ParseFiles("../The_Witcher/html/espèces.html"))
		fmt.Println("call server !!!!")
		tmpl.Execute(w, response)
	})

	http.ListenAndServe(":80", nil)
	fmt.Println("server closing")
}
