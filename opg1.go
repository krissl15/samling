package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"html/template"

)



func main() {
	http.HandleFunc("/1", path1)
	http.HandleFunc("/2", path2)
	http.HandleFunc("/3", path3)
	http.HandleFunc("/4", path4)
	http.ListenAndServe(":8080", nil)
}

//fields lol java4e

	type st1 []struct {
	Dato                string `json:"Dato"`
	Klokkeslett         string `json:"Klokkeslett"`
	Sted                string `json:"Sted"`
	Latitude            string `json:"Latitude"`
	Longitude           string `json:"Longitude"`
	AntallLedigePlasser string `json:"Antall_ledige_plasser"`
}


type st2 struct{
	Entries []struct{
 	East string `json:"east"`
 	North string	`json:"north"`
 	Zone string		`json:"zone"`
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
	}`json:"entries"`
}


	type st3 struct {
	Entries []struct {
	Latitude    string `json:"latitude"`
	Name        string `json:"name"`
	Adressenavn string `json:"adressenavn"`
	Longitude   string `json:"longitude"`
} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}


type st4 struct {
	Entries []struct {
		Name   string `json:"navn"`
		Number string `json:"nummer"`
	} `json:"entries"`
}

func path1(w http.ResponseWriter, r *http.Request) {

	page, err := http.Get("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json")
	if err != nil {
		log.Fatal(err)
	}
	jSonInfo, err := ioutil.ReadAll(page.Body)
	page.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var steder []st1
	er := json.Unmarshal(jSonInfo, &steder)
	if err != nil {
		fmt.Println("error:", er)
	}

	//fp := path.Join("templates", "sted.html")
	tmpl, err := template.ParseFiles("sted.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i:=0;i<len(steder);i++{
		if err := tmpl.Execute(w, steder); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}


}//slutt path1

func path2(w http.ResponseWriter, r *http.Request) {

	page, err := http.Get("https://hotell.difi.no/api/json/stavanger/lekeplasser?")
	if err != nil {
		log.Fatal(err)
	}
	jSonInfo, err := ioutil.ReadAll(page.Body)
	page.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var lekeplasser []st2
	er := json.Unmarshal(jSonInfo, &lekeplasser)
	if err != nil {
		fmt.Println("error:", er)
	}
	for i:=0;i<len(lekeplasser);i++ {
		fmt.Fprintf(w, "%+v", lekeplasser[i])

		fmt.Fprintln(w, "\n")
	}
}

func path3(w http.ResponseWriter, r *http.Request) {

	page, err := http.Get("https://hotell.difi.no/api/json/stavanger/utsiktspunkt?")
	if err != nil {
		log.Fatal(err)
	}
	jSonInfo, err := ioutil.ReadAll(page.Body)
	page.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var punkt []st3
	er := json.Unmarshal(jSonInfo, &punkt)
	if err != nil {
		fmt.Println("error:", er)
	}
	for i:=0;i<len(punkt);i++ {
		fmt.Fprintf(w, "%+v", punkt[i])
		fmt.Fprintln(w, "\n")
	}
}

func path4(w http.ResponseWriter, r *http.Request) {

	page, err := http.Get("https://hotell.difi.no/api/json/difi/geo/fylke")
	if err != nil {
		log.Fatal(err)
	}
	jSonInfo, err := ioutil.ReadAll(page.Body)
	page.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var punkt []st4
	er := json.Unmarshal(jSonInfo, &punkt)
	if err != nil {
		fmt.Println("error:", er)
	}
	for i:=0;i<len(punkt);i++ {
		fmt.Fprintf(w, "%+v", punkt[i])
		fmt.Fprintln(w, "\n")
	}
}

func foo5(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nummer5"))
}


