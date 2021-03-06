package database

import (
	"Avito/Application/io_data"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const delete_all = "DELETE FROM statistic";
const _select = "SELECT * FROM statistic WHERE date between ? AND ?";
const insert = "INSERT INTO statistic VALUES(?,?,?,?)";

//не очень понял про опциональные поля в условии, понял так что если придет запрос с такой же датой,
//то нужно добавить clicks и views к уже имеющимся, а cost обновить
const update = "UPDATE statistic SET views = views + ?, clicks = clicks + ?, cost = ?  WHERE date = ?"
const select_date = "SELECT * FROM statistic WHERE date = ?";

type 	Store struct {
	db  *sql.DB;
	conf io_data.Config;
}

func InitDataSourceName(conf io_data.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", conf.User, conf.Pass, conf.Host, conf.Dbname);
}

func New(conf io_data.Config) *Store {
	return &Store{conf: conf};
}

func OpenConnection(st *Store) error {
	db, err := sql.Open("mysql", InitDataSourceName(st.conf));
	if (err != nil) {
		log.Println("Error open connection with database")
		return errors.New("mysql error")
	}
	st.db = db;
	return nil;
}

func CloseConnection(st *Store)  {
	st.db.Close();
}

func Select(st *Store, from string, to string, arr *[]*Data) error {
	rows, err := st.db.Query(_select, from, to);

	if err != nil {
		log.Println("Error SELECT from database")
		return errors.New("mysql error")
	}
	for (rows.Next()) {
		info := new(Data);
		date := new(time.Time);
		err := rows.Scan(&date, &info.Views, &info.Clicks, &info.Cost);
		if (err != nil) {
			return errors.New("mysql error");
		}
		//дату сохраним в форматированной строке для удобства отдачи пользователю
		info.Date = date.Format("2006-01-02");
		addData(info);
		*arr = append(*arr, info);
		}
		return nil;
}

func Delete(st *Store) error {
	_, err := st.db.Exec(delete_all);
	if (err != nil){
		return errors.New("mysql error");
	}
	return nil;
}

//будем сразу добавлять  cpc и cpm, чтобы не проходить по массиву еще раз перед сериализацией в JSON
func  	addData(data *Data) {
	if data.Clicks == 0 {
		data.Cpc = -1;
	} else {
		data.Cpc = data.Cost / float32(data.Clicks);
	}
	if data.Views == 0 {
		data.Cpm = -1
	} else {
		data.Cpm = (data.Cost / float32((data.Views))) * 1000
	}
}

func 	PostManager(data *io_data.Json, st *Store) error  {
	rows, err := st.db.Query(select_date, data.Date);
	if err != nil {
		log.Println("Error SELECT from database")
		return errors.New("mysql error")
	}
	if (!rows.Next()) {
		_, err := st.db.Exec(insert, data.Date, data.Views, data.Clicks, data.Cost);
		if err != nil {
			log.Println("Error INSERT MySQL");
			return errors.New("mysql error");
		}
	} else {
		_, err := st.db.Exec(update, data.Views, data.Clicks, data.Cost, data.Date);
		if err != nil {
			log.Println("Error UPDATE MySQL");
			return errors.New("mysql error");
		}
	}
	return nil;
}


