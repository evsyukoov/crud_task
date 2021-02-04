package app

import (
	"Avito/Application/database"
	"Avito/Application/io_data"
	"Avito/Application/methods"
	"fmt"
	"log"
	"net/http"
)

func sendCodeToClient(w http.ResponseWriter, code int)  {
	w.WriteHeader(code);
}

func sendAnswerToGet(w http.ResponseWriter, body []byte)  {
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(body)));
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(200);
	w.Write(body);
}

func errorHandler(w http.ResponseWriter, err error)  {
	if err != nil && err.Error() == "bad request" {
		sendCodeToClient(w, 400);
	}
	if err != nil && err.Error() == "mysql error" {
		sendCodeToClient(w, 500);
	}
}

func Handler() http.HandlerFunc {
	var store *database.Store
	err, conf := io_data.InitDbParametrs();
	if err != nil {
		log.Fatal("Error in config file..Stop!")
	} else {
		store = database.New(*conf);
	}
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			body, err := methods.Get(*r, store);
			errorHandler(w, err)
			if body != nil {
				sendAnswerToGet(w, body);
			}
		} else if r.Method == http.MethodDelete {
			err := methods.Delete(*r, store)
			errorHandler(w, err);
			if err == nil {
				sendCodeToClient(w, 204);
			}
		} else if r.Method == http.MethodPost {
			err := methods.Post(*r, store)
			errorHandler(w, err);
			if err == nil {
				sendCodeToClient(w, 200)
			}
		} else {
			sendCodeToClient(w, 400);
		}
	}
}
