package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver
)

// connection
func CreateConnection() *sql.DB {
	//load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error Loading .env file")
	}

	//DB Connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	//check connection DB
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("DB Connection Success!")
	return db
}

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}
