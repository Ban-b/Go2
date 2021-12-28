package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type AppSetting struct {
	appHost string
	appPort string
}

func (h *AppSetting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Server running on address: %s:%s", h.appHost, h.appPort)
	w.Write([]byte(msg))
}

\\Создание паники

func main() {

	appSettings := AppSetting{appHost: "localhost", appPort: "8080"}

	defer func() {
		for {
			log.Printf("trying start server on %s:%s address \n", appSettings.appHost, appSettings.appPort)
			err := http.ListenAndServe(appSettings.appHost+":"+appSettings.appPort, nil)
			if err != nil {
				ErrInvalidPort(appSettings.appPort, appSettings.appPort, err)

				newPort, _ := strconv.Atoi(appSettings.appPort)
				newPort += 10
				appSettings.appPort = strconv.Itoa(newPort)
			}
		}
	}()

	http.Handle("/", &appSettings)
}

\\Собственная ошибка.  

func ErrInvalidPort(appHost string, appPort string, err error) {
	log.Printf("address %s:%s in use\nError:%s", appHost, appPort, err)
}

\\Вывод её на консоль.
