package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mcgtrt/go-puerto/api"
	"github.com/mcgtrt/go-puerto/storage"
	"github.com/mcgtrt/go-puerto/utils"
)

func main() {
	config, err := utils.NewDefaultConfig()
	if err != nil {
		panic("configuration error: " + err.Error())
	}
	store, err := storage.NewStore(config)
	if err != nil {
		panic("store initialisation error:" + err.Error())
	}
	handler := api.NewHandler(store, config.HTTP)
	router := api.NewRouter(handler)

	fmt.Println("http server running on port", config.HTTP.Port)
	http.ListenAndServe(":"+strconv.Itoa(config.HTTP.Port), router)
}
