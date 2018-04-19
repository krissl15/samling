package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func main() {
	http.HandleFunc("/1", path1)
	http.HandleFunc("/2", path2)
	http.HandleFunc("/3", foo3)
	http.HandleFunc("/4", foo4)
	http.ListenAndServe(":8080", nil)
}

//fields lol java4e

type st1 struct {
	Sted string `json:"Sted"`
	Dato  string `json:"Dato"`
	Klokkeslett string `json:"Klokkeslett"`
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
	Antall_ledige_plasser string `json:"Antall_ledige_plasser"`
}

type st2 struct{
 	East string `json:"navn"`
 	North string	`json:"nummber"`
 	Zone string		`json:"zone"`
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
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
	for i:=0;i<len(steder);i++ {
		fmt.Fprintf(w, "%+v", steder[i])
		fmt.Fprintln(w, "\n")
	}

}

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

func foo3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nummer3"))
}

func foo4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nummer34"))
}

func foo5(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nummer5"))
}

func dataSet1(){


}

