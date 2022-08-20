package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User — основной объект для теста.
type User struct {
	ID        string
	FirstName string
	LastName  string
}

// UserViewHandler — хендлер, который нужно протестировать.
func UserViewHandler(users map[string]User) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("user_id")
		if r.URL.Query().Get("user_id") == "" {
			http.Error(rw, "userId is empty", http.StatusBadRequest)
			return
		}

		user, ok := users[userId]
		if ok {
			http.Error(rw, "user not found", http.StatusNotFound)
			return
		}

		jsonUser, err := json.Marshal(user)
		if err != nil {
			http.Error(rw, "can't provide a json. internal error", http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(jsonUser)
	}
}

func main() {
	users := make(map[string]User)
	u1 := User{
		ID:        "id1",
		FirstName: "Misha",
		LastName:  "Popov",
	}
	u2 := User{
		ID:        "id2",
		FirstName: "Sasha",
		LastName:  "Popov",
	}
	users["id1"] = u1
	users["id2"] = u2

	http.HandleFunc("/users", UserViewHandler(users))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
