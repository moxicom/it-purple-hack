package api_handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/moxicom/it-purple-hack/internal/services/price"
	"github.com/moxicom/it-purple-hack/models"
)

func get_price(price_service *price.PriceService, storage models.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		r.ParseForm()
		decoder := json.NewDecoder(r.Body)
		var data models.PriceRequest
		err := decoder.Decode(&data)
		if err != nil {
			log.Fatalln("Invalid Json " + err.Error())
		}
		price, err := price_service.GetPrice(data, storage)
		if err != nil {
			log.Fatalln("Invalid GetPrice func " + err.Error())
		}
		err = json.NewEncoder(w).Encode(price)
		if err != nil {
			log.Fatalln("Bad response " + err.Error())
		}
	}
}

func update_price(price_service *price.PriceService, storage models.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		r.ParseForm()
		decoder := json.NewDecoder(r.Body)
		var data models.UpdateRequest
		err := decoder.Decode(&data)
		if err != nil {
			log.Fatalln("Invalid Json " + err.Error())
		}
		err = price_service.UpdatePrice(data, storage)
		if err != nil {
			log.Fatalln("Invalid GetPrice func " + err.Error())
		}
		err = json.NewEncoder(w).Encode("{status:ok}")
		if err != nil {
			log.Fatalln("Bad response " + err.Error())
		}
	}
}

func delete_price(price_service *price.PriceService, storage models.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		r.ParseForm()
		decoder := json.NewDecoder(r.Body)
		var data models.UpdateRequest
		err := decoder.Decode(&data)
		if err != nil {
			log.Fatalln("Invalid Json " + err.Error())
		}
		err = price_service.DeleteDiscount(data, storage)
		if err != nil {
			log.Fatalln("Invalid GetPrice func " + err.Error())
		}
		err = json.NewEncoder(w).Encode("{status:ok}")
		if err != nil {
			log.Fatalln("Bad response " + err.Error())
		}
	}
}
