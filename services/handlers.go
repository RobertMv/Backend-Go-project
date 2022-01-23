package services

import (
	"awesomeProject/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "300P@midorov"
	dbName   = "go_restaurants"
)

type response struct {
	ID      int64
	Message string
}

func createConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database.")
	return db
}

func CreatePosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var position models.Position
	err := json.NewDecoder(r.Body).Decode(&position)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := insertPosition(position)
	res := response{insertID, "User created successfully."}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatalf("Unable to encode the request body. %v", err)
	}
}

func GetPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	position, err := getUser(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}
	err = json.NewEncoder(w).Encode(position)
	if err != nil {
		log.Fatalf("Unable to encode position.")
	}
}

func UpdatePosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.")
	}

	var position models.Position
	err = json.NewDecoder(r.Body).Decode(&position)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updateRows := updatePosition(int64(id), position)
	res := response{int64(id), fmt.Sprintf("Position updated successfully. Total rows/records affected %v", updateRows)}

	json.NewEncoder(w).Encode(res)
}

func DeletePosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	deletedRows := deletePosition(int64(id))
	res := response{int64(id), fmt.Sprintf("Position deleted successfully. Total rows/records affected %v", deletedRows)}
	json.NewEncoder(w).Encode(res)
}

func getUser(id int64) (models.Position, error) {
	db := createConnection()
	defer db.Close()

	var position models.Position
	sqlStatement := "SELECT * FROM positions WHERE id=$1"

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&position.Id, &position.Code, &position.Name, &position.Salary)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return position, nil
	case nil:
		return position, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return position, err
}

func insertPosition(position models.Position) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := "INSERT INTO positions (code, name, salary) VALUES ($1, $2, $3)"
	var id int64

	err := db.QueryRow(sqlStatement, position.Code, position.Name, position.Salary)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)
	return id
}

func updatePosition(id int64, position models.Position) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := "UPDATE positions SET code=$2, name=$3, salary=$4 WHERE id=$1"
	res, err := db.Exec(sqlStatement, id, position.Code, position.Name, position.Salary)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}

func deletePosition(id int64) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := "DELETE FROM positions WHERE id=$1"
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}
