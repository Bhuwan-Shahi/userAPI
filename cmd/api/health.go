package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("okay everyting is fine!!! and this is live reloading\n"))
}
