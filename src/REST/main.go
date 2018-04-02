package main

import (

	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Box struct {
	ID string `json:"id,omitempty"`
	X string `json:"x,omitempty"`
	Y string `json:"y,omitempty"`
	Z string  `json:"z,omitempty"`
}


var Boxes []Box

func GetPositionBox(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for _, item := range Boxes {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Box{})
}

func GetBoxes(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(Boxes)
}

func CreateBox(w http.ResponseWriter, req *http.Request){
	req.ParseForm()
	var box Box
	byteBox, _ := json.Marshal(req.Form)
	for i := 0; i < len(byteBox) ; i++ {
		if byteBox[i] == 93 || byteBox[i] == 91 {
			byteBox = append(byteBox[:i], byteBox[i+1:]...)
		}
	}
	err := json.Unmarshal(byteBox, &box)
	if err != nil {
		panic(err)
	}
	Boxes = append(Boxes, box)
	json.NewEncoder(w).Encode(Boxes)
}

func DeleteBox(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for index, item := range Boxes {
		if item.ID == params["id"] {
			Boxes = append(Boxes[:index], Boxes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Boxes)
}

func ChangePositionBox(w http.ResponseWriter, req *http.Request){
	req.ParseForm()
	var box Box
	byteBox, _ := json.Marshal(req.Form)
	for i := 0; i < len(byteBox) ; i++ {
		if byteBox[i] == 93 || byteBox[i] == 91 {
			byteBox = append(byteBox[:i], byteBox[i+1:]...)
		}
	}
	err := json.Unmarshal(byteBox, &box)
	if err != nil {
		panic(err)
	}

	for index, item := range Boxes {
		if item.ID == box.ID {
			item.X = box.X
			item.Y = box.Y
			item.Z = box.Z
			Boxes = append(Boxes[:index], Boxes[index+1:]...)
			Boxes = append(Boxes, box)
			log.Println(Boxes)
			break
		}
	}
	json.NewEncoder(w).Encode(Boxes)
}


func main(){
	router := mux.NewRouter()

	//people = append(people, Person{ID: "1", Firstname: "Nick", Lastname: "Ber", Address: &Address{City: "Portland", State: "USA"}})
	//people = append(people, Person{ID: "2", Firstname: "Addalinda", Lastname: "Sheid"})
	router.HandleFunc("/boxes", GetBoxes).Methods("GET")
	router.HandleFunc("/boxes/{id}", GetPositionBox).Methods("GET")
	router.HandleFunc("/boxes/{id}/create", CreateBox).Methods("POST")
	router.HandleFunc("/boxes/{id}", DeleteBox).Methods("DELETE")
	router.HandleFunc("/boxes/{id}", ChangePositionBox).Methods("POST")


	log.Fatal(http.ListenAndServe(":12345", router))
}
