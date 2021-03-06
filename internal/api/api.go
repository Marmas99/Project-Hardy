package api

import (
	"io/ioutil"
	"net/http"
	"project/hardy/internal/database"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var apilogger *logrus.Logger

func StartServer(logger *logrus.Logger) {
	myRouter := mux.NewRouter().StrictSlash(true)
	apilogger = logger

	apilogger.Infoln("Initiating handlers")
	registerhandlers(myRouter)

	apilogger.Infoln("API listening...")
	http.ListenAndServe(":3030", myRouter)
}

func registerhandlers(router *mux.Router) {
	router.HandleFunc("/insert", insertItem).Methods("POST")
	router.HandleFunc("/delete", deleteItem).Methods("POST")
}

func insertItem(writer http.ResponseWriter, request *http.Request) {
	reqData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		apilogger.Errorln("Error while reading request body")
		panic("Stopping program")
	}
	database.InsertNewItem(reqData, apilogger)
}

func deleteItem(writer http.ResponseWriter, request *http.Request) {
	reqData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		apilogger.Errorln("Error while reading request body")
		panic("Stopping program")
	}
	database.DeleteItem(reqData, apilogger)
}
