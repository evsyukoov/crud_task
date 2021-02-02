package methods

import (
	"AvitoTest/database"
	"AvitoTest/io_data"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func checkSortedField(field string) bool {
	if field != "views" && field != "clicks" && field != "cost" && field != "date" {
		return false;
	}
	return true;
}

func parsePath(path string, req *io_data.Request) error  {
	var route []string;
	route = strings.Split(path, "/");
	if (route[1] != "show") {
		return errors.New("bad request");
	}
	if len(route) > 3 || len(route) == 3 && !checkSortedField(route[2]) {
			return errors.New("bad request");
	}
	if (len(route) == 3) {
		req.Sort_field = route[2];
	} else {
		req.Sort_field = "date";
	}
	return nil;
}

func CheckDate(format, date string) bool {
	_, err := time.Parse(format, date)
	if err != nil {
		return false
	}
	return true
}

func 	parseQuery(query url.Values, req *io_data.Request) error {
	if len(query) > 2 || query.Get("from") == "" || query.Get("to") == "" ||
		!CheckDate("2006-01-02", query.Get("from")) || !CheckDate("2006-01-02", query.Get("to")) {
		return errors.New("bad request");
	}
	req.From = query.Get("from");
	req.To = query.Get("to");
	return nil;
}

  //  GET /show/cost?from=2020-01-11&to=2021-02-30 HTTP/1.1
func parseGet(r http.Request, req *io_data.Request) error {
	if err := parsePath(r.URL.Path, req); err != nil {
		log.Println("error while parsing GET route")
		return err;
	}
	if err := parseQuery(r.URL.Query(), req); err != nil {
		log.Println("Error while parsing GET query")
		return err;
		}
	return nil;
}

// наверх пробрасываем 2 типа ошибок
// 1) Bad request в случае неравильного пользовательского ввода
// 2) Internal server error в случае проблем с подключением/выборкой из БД

func Get(r http.Request, store *database.Store) ([]byte, error)   {
	req := new(io_data.Request);
	arr := make([]*database.Data, 0);
	if err := parseGet(r, req); err != nil {
		return nil, err;
	}
	if err := database.OpenConnection(store); err != nil {
		return nil, err;
	}
	err := database.Select(store, req.From, req.To, &arr);
	if (err != nil) {
		database.CloseConnection(store);
		return nil, err;
	}
	database.CloseConnection(store);
	database.PrintArr(arr);
	database.SortFromTypes(*req, arr);
	json_data, _ := json.Marshal(arr);
	return json_data, nil;
}

