package methods

import (
	"AvitoTest/database"
	"AvitoTest/io_data"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)


func parsePostRoute(path string) error {

	if path != "/save" {
		return errors.New("bad request")
	}
	return nil;
}

func checkInt(num float64) bool  {
	s := fmt.Sprintf("%f", num);
	matcher, _ := regexp.MatchString("\\d*\\.[0]*$",s)
	return matcher;
}

func	parseMap(f interface{}, obj *io_data.Json) (error) {
	err := errors.New("bad request");
	m := f.(map[string]interface{});
	for key, val := range m {
		switch v := val.(type) {
		case string:
			if key != "date" {
				log.Println("Unknown key in JSON body")
				return err;
			}
			if  !CheckDate("2006-01-02", v) {
				log.Println("Wrong date format in JSON");
				return errors.New("bad request");
			}else {
				obj.Date = v;
				obj.IsContainsDate = true;
			}
			break
		case float64:
			if (v < 0) {
				log.Println("Error while parsing JSON. Number can't be a negative");
				return err;
			}
			if key != "cost" && !checkInt(v) {
				log.Println("Error while parsing JSON. Number of views or clicks can't be a float");
				return err;
			}
			if key == "views" {
				obj.Views = int(v);
			} else if key == "clicks" {
				obj.Clicks = int(v);
			} else if key == "cost" {
				obj.Cost = float32(v);
			} else {
				log.Println("Error while parsing JSON. Unknown key")
				return err;
			}
			break
		default:
			log.Println("Unknown key type");
			return err
		}
	}
	if !obj.IsContainsDate {
		log.Println("Not find mandatory key date");
		return err;
	}
	return nil;
}

func decodeJSON(body[] byte, obj *io_data.Json)   error{
	var f interface{};
	if json.Unmarshal(body, &f) != nil {
		log.Println("Bad JSON body");
		return errors.New("bad request");
	}
	if err := parseMap(f, obj); err != nil {
		return err;
	}
	return nil;
}

func Post(r http.Request, store *database.Store) error  {
	if err := parsePostRoute(r.RequestURI); err != nil {
		return err;
	}
	body, err := ioutil.ReadAll(r.Body);
	if err != nil {
		log.Println("Can't read body")
		return errors.New("io problems");
	}
	obj := io_data.NewJson();
	if err = decodeJSON(body, obj); err != nil {
		return err;
	}
	if err := database.OpenConnection(store); err != nil {
		return err;
	}
	err = database.PostManager(obj, store);
	if (err != nil) {
		database.CloseConnection(store);
		return err;
	}
	database.CloseConnection(store);
	return nil;

}
