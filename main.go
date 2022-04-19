package main


import(
	"net/http"
	"encoding/json"
	"time"
	"fmt"
	"log"
	"github.com/gorilla/mux"

)

type CurrentDate struct{
	CurrentDate time.Time
}

type User struct {
	Username string
	Password string `json:"-"`
	IsAdmin bool
	CreatedAt time.Time
}



func main(){
	log.Println("starting API server")
	router := mux.NewRouter()
	log.Println("creating routes")
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/date", currentDate).Methods("GET")
	router.HandleFunc("/print", echoHandler).Methods("POST")
	
	
	http.ListenAndServe(":3000", router)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func currentDate(w http.ResponseWriter, r *http.Request){
	date := CurrentDate{}
	date.CurrentDate = time.Now().Local()

	currentDateJson, err := json.Marshal(date)
	if err != nil{
		panic(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK) 
	w.Write(currentDateJson)
}


func echoHandler(w http.ResponseWriter, r *http.Request){
	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		panic(err)
	}

	user.CreatedAt = time.Now().Local()

	userJson, err := json.Marshal(user)
	if err != nil{
		panic(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK) 
	w.Write(userJson)
}

