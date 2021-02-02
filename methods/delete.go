package methods

import (
	"AvitoTest/database"
	"errors"
	"net/http"
)

//DELETE /clear
func parseDelete(path string) error  {

	if path != "/clear" {
		return errors.New("bad request")
	}
	return nil;
}

func Delete(r http.Request, store *database.Store) error  {
	//fmt.Printf("eto on %s" ,r);
	err := parseDelete(r.RequestURI)
	if err != nil {
		return err;
	}
	if err := database.OpenConnection(store); err != nil {
		return err;
	}
	err = database.Delete(store);
	if err != nil {
		database.CloseConnection(store);
		return err;
	}
	database.CloseConnection(store);
	return nil;
}
