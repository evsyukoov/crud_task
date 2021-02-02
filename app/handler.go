package app

import (
	"AvitoTest/database"
	"AvitoTest/methods"
	"fmt"
	"net/http"
)

//метод показа статистики
// GET /show?from=2020-11-01&to=2012-01-02 - по умолчанию по дате
// GET /show/clicks?from=2020-11-01&to=2012-01-02  - сортировка по определенному полю (доп.)

//метод сохранения стаистики
//	POST /save
//
//	{JSON}

// метод удаления статистики
//DELETE /clear - удаление всего

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

func Handler(w http.ResponseWriter, r *http.Request) {
	store := database.New();
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
