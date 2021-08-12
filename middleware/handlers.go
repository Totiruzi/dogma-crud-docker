package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Totiruzi/dogma-crud-docker/models"
	"github.com/astaxie/beego/context/param"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

// response format
type response struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:message,omitempty"`
}

// createConnection creates connection with postgres db
func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()


	if err != nil {
		panic(err)
	}

	fmt.Print("Successfully connected!")
	// return the connection
	return db
}

// CreateUser create a user in the postgres db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty user of type models.User
	var user models.User

	// decode the json request to user
	err :=json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal("Unabe to decode the request body.%v", err)
	}

	// call insert user function and pass the user
	insertID := insertUser(user)

	// format a response object
	res := response {
		ID: insertID,
		Message: "User created ccessfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetUser will return a single user by its id
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the userid from he equest params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal("Unable to convert the string to int. %v", err)
	}

	// call the getUser function with user id to retrieve a single user
	user, err := getUser(int64(id))

	if err != nil {
		log.Fatalf("Unable to get ser. %v, err")
	}

	// send he esponse
	json.NewEncoder(w).Encode(user)
}

// GetAllUser will return all the users
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "appliation/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the users in the db
	users, err := getAllUser()
	
	if err != nil {
		log.Fatalf("Unable to get all users. %v", err)
	}

	// send all the users as a response
	json.NewEncoder(w).Encode(users)
} 

// UpdateUser update user's detail in the postgres db
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//get the userId from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from sring to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal("Unable to convert the string into int. %v, err")
	}

	// create an empty user of type model.User
	var user model.User

	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v, err")
	}

	// call the update user to update the user
	updateRows := updateUser(int64(id), user)

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updateRows)

	//format the response message
	res := response {
		ID: int64(id)
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteUser delete user's detail in the postgres db
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userId from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in sring to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	//call the deleteUser, convert the int to int64
	deletedRow := deleteUser(int64(id)
	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/records affected %v", deletedRows)

	// format the esponse message
	res := response {
		ID: int64(id)
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// InsertUser insert's one user into the DB
func insertUser(user models.User) int64 {
	// create the posgres db connection
	db := createConnection()

	close the db connection
	defer db.close()

	// create the insert sql query
	// returning userid will return the id of the inserted user 
	w.Header().Set("Access-Control-Allow-")
	sqlStatement := `INSERT INTO 	users(name, location, age) VALUES ($, $2, $3) RETURNING userid`

	// the inserted id will be stored in this id
	var id int64

	// execute the sql statement
	// can function will save the inserted id into the id
	err :=db.QueryRow(sqlStatemet, user.Name, user.Location, user.Age).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

// getUser get one user from the DB by its userid
func getUser(id int64) (models.User, error) {
	// create the postgres data connection
	db := createConnection()

	// close the db connection
	defer db.close()

	// create a user of models.User type
	var user models.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE userid=$1`

	// execute the statement
	row = db.QueryRow(&user.ID, &user.Name, &user.Age, &user.Location)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows where returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)		
	}
	// return empty user on error
	return user, err
}

// get one user from the DB by it's userid
func getAllUsers() ([]models.User, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var users []models.User

	// create the select statement
	sqlStatement := `SELECT * FROM users`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user models.User

		// unmarshal the row object to user
		err = rows.Scan($user.ID, &user.Name, &user.Age, &user.Location)

		if err != nil {
			log.Fatalf("Unable to scan row. %v", err)
		}

		// append the user in the user slice
		users = append(users, user)
	}

	// return empty user on error
	return users, err
}

// updateUser update the user in the database
func updateUser(id int64, user models.User) int64 [
	// create the postgres db connection
	db := createConnection()

	// close the db connetion
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, user.Location, user.Age)

	if err != nil {
		log.Fatalf("Un able to execute query. %v", err)
	}

	// check how many rows where affected
	rowsAffected, err := res.rowsAffected()

	if err != nil  {
		log.Fatalf("Error while checking affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
]

