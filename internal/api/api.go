package api

import (
	"io/ioutil"
	"net/http"
	"project/hardy/internal/item"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var apilogger *logrus.Logger

func StartServer(logger *logrus.Logger) {
	myRouter := mux.NewRouter().StrictSlash(true)
	apilogger = logger

	apilogger.Info("Initiating handlers")
	registerhandlers(myRouter)

	apilogger.Info("Server listening...")
	http.ListenAndServe(":3030", myRouter)
}

func registerhandlers(router *mux.Router) {
	router.HandleFunc("/insert", insertItem).Methods("POST")
}

func insertItem(writer http.ResponseWriter, request *http.Request) {
	reqData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		apilogger.Errorln("Error while reading request body")
		panic("Stopping program")
	}
	item.CreateNewItem(reqData, writer)
}
