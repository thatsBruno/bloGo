package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	repondWithError(w, 400, "Somethig's wrong")
}
