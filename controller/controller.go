package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"anshulgithub.com/anshul/usermangement/database"
	"anshulgithub.com/anshul/usermangement/helper"
	"anshulgithub.com/anshul/usermangement/models"
	"github.com/gorilla/mux"
)

var uniqueUserId int

func GetContrller() mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/users", getAllUser).Methods("GET")
	router.HandleFunc("/users/{id}", getOneUser).Methods("GET")
	router.HandleFunc("/users", addUser).Methods("POST")
	router.HandleFunc("/users", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", removeUser).Methods("DELETE")
	router.Handle("/", router)
	return *router

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to home page")
}
func getAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetAllUser()
	if err != nil {
		w.WriteHeader(http.StatusOK)
	}
	helper.ErrCheck(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func getOneUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// use the ParseInt() Function
	userid, _ := strconv.ParseInt(id, 10, 64)
	user, err := database.ReadUser(int(userid))
	helper.ErrCheck(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func addUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	if user.UserId == 0 {
		rand.Seed(time.Now().UnixNano())

		// Generate a random number
		uniqueRandomUserId := rand.Intn(1000000)
		uniqueUserId = uniqueRandomUserId
		user.UserId = uniqueRandomUserId
	} else {
		uniqueUserId = user.UserId
	}
	// the validation form email id
	if !helper.IsValidEmail(user.GmailId) {
		resp := make(map[string]string)
		resp["message"] = "Please give the valid Email"
		jsonResp, _ := json.Marshal(resp)
		//helper.ErrCheck(err)
		w.Write(jsonResp)
	}
	if !helper.IsValidNumber(user.PhoneNo) {
		resp := make(map[string]string)
		resp["message"] = "Please give the valid number"
		jsonResp, _ := json.Marshal(resp)
		//helper.ErrCheck(err)
		w.Write(jsonResp)
	}
	if helper.EmtRequest(&user) {
		resp := make(map[string]string)
		resp["message"] = "User name , User address and User Phone number is required"
		jsonResp, _ := json.Marshal(resp)
		//helper.ErrCheck(err)
		w.Write(jsonResp)
	}
	if !helper.EmtRequest(&user) && helper.IsValidEmail(user.GmailId) && helper.IsValidNumber(user.PhoneNo) {
		err := database.CreateUser(user)
		//helper.ErrCheck(err)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := make(map[string]string)
			resp["userid"] = strconv.Itoa(uniqueUserId)
			resp["message"] = err.Error()
			jsonResp, err := json.Marshal(resp)
			helper.ErrCheck(err)
			w.Write(jsonResp)
		} else {
			w.WriteHeader(http.StatusCreated)
			resp := make(map[string]string)
			resp["userid"] = strconv.Itoa(uniqueUserId)
			resp["message"] = "User is created sucessfully"
			jsonResp, err := json.Marshal(resp)
			helper.ErrCheck(err)
			w.Write(jsonResp)
		}
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	w.Header().Set("Content-Type", "application/json")
	if helper.EmtRequest(&user) {
		resp := make(map[string]string)
		resp["message"] = "User name , User address and User Phone number is required for updating user"
		jsonResp, err := json.Marshal(resp)
		helper.ErrCheck(err)
		w.Write(jsonResp)
	} else {
		err := database.UpdateUser(&user)
		//helper.ErrCheck(err)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := make(map[string]string)
			resp["message"] = err.Error()
			jsonResp, err := json.Marshal(resp)
			helper.ErrCheck(err)
			w.Write(jsonResp)
		} else {
			w.WriteHeader(http.StatusAccepted)
			resp := make(map[string]string)
			resp["message"] = "User is updated sucessfully"
			jsonResp, err := json.Marshal(resp)
			helper.ErrCheck(err)
			w.Write(jsonResp)
		}
	}
}

func removeUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	userid, _ := strconv.ParseInt(id, 10, 64)
	err := database.DeleteUser(int(userid))
	//helper.ErrCheck(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := make(map[string]string)
		resp["message"] = err.Error()
		jsonResp, err := json.Marshal(resp)
		helper.ErrCheck(err)
		w.Write(jsonResp)
	} else {
		w.WriteHeader(http.StatusAccepted)
		resp := make(map[string]string)
		resp["message"] = "User is updated sucessfully"
		jsonResp, err := json.Marshal(resp)
		helper.ErrCheck(err)
		w.Write(jsonResp)
	}

}
