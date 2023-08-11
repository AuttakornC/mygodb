package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Mygodb struct {
	db *sql.DB
}

type Config struct {
	Host string `default:"localhost"`
	Port string `default:"5432"`
	User string
	Pass string
	DBnm string
}

func (m *Mygodb) Connect(config Config) error {
	parsed_config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Pass, config.DBnm, config.Port)
	db, err := sql.Open("postgres", parsed_config)

	// data, erro := db.Query("SELECT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND	 tablename  = 'student');")
	// if erro != nil {
	// 	fmt.Println(erro)
	// } else {
	// 	for data.Next() {
	// 		var a bool
	// 		if err := data.Scan(&a); err != nil {
	// 			fmt.Println(err)
	// 		} else {
	// 			fmt.Println(a)
	// 		}

	// 	}
	// }
	if err != nil {
		fmt.Println("+-----------------------------------------------------+")
		fmt.Println("| [ MyGoDB ] : Cannot connect to the Database!!       |")
		fmt.Println("+-----------------------------------------------------+")
		fmt.Println(err)
		return err
	} else {
		fmt.Println("+-----------------------------------------------------+")
		fmt.Println("| [ MyGoDB ] : Connect to the Database Successfully!! |")
		fmt.Println("+-----------------------------------------------------+")
		m.db = db
		return nil
	}
}
