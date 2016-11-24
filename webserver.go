package main

import
(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	_"github.com/lib/pq"
	"database/sql"
)

var database *sql.DB

type Users struct {
  Users []User `json:"users"`
}

type User struct {
  ID int "json:id"
  Name  string "json:username"
  Email string "json:email"
  First string "json:first"
  Last  string "json:last"
}

func main() {

  db, err := sql.Open("postgres", "user=manuel dbname=golang_webservices sslmode=disable")
  if err != nil {
    panic(err.Error())
  }
  database = db
  routes := mux.NewRouter()
  routes.HandleFunc("/api/users", UserCreate).Methods("POST")
  routes.HandleFunc("/api/users", UsersRetrieve).Methods("GET")
  http.Handle("/", routes)
  http.ListenAndServe(":8080", nil)
}
func UserCreate(w http.ResponseWriter, r *http.Request) {
  NewUser := User{}
  NewUser.Name = r.FormValue("user")
  NewUser.Email = r.FormValue("email")
  NewUser.First = r.FormValue("first")
  NewUser.Last = r.FormValue("last")
  output, err := json.Marshal(NewUser)
  fmt.Println(string(output))
  if err != nil {
    fmt.Println("Something went wrong!")
  }

  rows, err := database.Query("INSERT INTO users (user_nickname, user_first, user_last, username, user_email) VALUES (NewUser.Name, NewUser.First, NewUser.Last, NewUser.Email)")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(rows)
}



func UsersRetrieve(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Pragma","no-cache")
  rows, err := database.Query("SELECT * FROM users")
  if err != nil{
    fmt.Println(err)
  }

  Response 	:= Users{}

  for rows.Next() {
    
    user := User{}

    err := rows.Scan(&user.Name,&user.First, &user.Last, &user.Email, &user.ID)
    if err != nil{
      panic(err.Error())
    }
    
  Response.Users = append(Response.Users, user)
  }
   output,_ := json.Marshal(Response)
  fmt.Fprintln(w,string(output))
}

