package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"native-go-api/models"
	"net/http"
	"strconv"
	"time"
)

// Articles ...
var Articles []models.Article

var (
	Promotions []models.Promotion
)

func init() {

}

func main() {

	// Setup and start server
	port := ":8000"
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/promotions", returnAllPromotions)
	router.HandleFunc("/promotions/createPromotion", createPromotions).Methods("POST")
	server := &http.Server{
		Handler: router,
		Addr:    port,
		// timeouts so the server never waits forever...
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServe())

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func returnAllPromotions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Promotions)
}

func createPromotions(res http.ResponseWriter, req *http.Request) {
	// set the header "Content-Type: application/json"
	res.Header().Set("Content-Type", "application/json")
	// assign a promotion variable of type Promotion struct
	var promotion models.Promotion
	if err := json.NewDecoder(req.Body).Decode(&promotion); err != nil {
		log.Println("ERROR: " + err.Error())
		return
	}

	promotion.Id = strconv.Itoa(rand.Intn(1000000)) // just creating dummy id as we are not sending an id from postman - not safe for production
	// Now, we will simply append that course with decoded req.Body in our
	// courses slice (Array) which we have defined] earlier
	Promotions = append(Promotions, promotion)
	// At the end, send the course that we have created.
	json.NewEncoder(res).Encode(promotion)
}

/*func createPromotions(promotion models.Promotion) error {
	bd, err := db.getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO PROMOTION (ID, PRODUCT_CODE, BOOK) VALUES (?, ?)", promotion.ProductCode, promotion.ISBN)
	return err
}*/
