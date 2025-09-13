package main

import (
	"net/http"
)

func handlerReadiness(dw http.ResponseWriter, code int, payload interface{}){
	respondWithJSON(dw,200,struct{}{})
}
