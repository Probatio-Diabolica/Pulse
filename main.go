package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)


func main(){

	//loads the env automatically
	err := godotenv.Load()

	if err != nil{
		log.Fatal("cannnot load .env file")
	}

	portString:= os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port was not found in the environment")
	}

	router := chi.NewRouter()
	
	router.Use(cors.Options{
		AllowedOrigins 		: []string{"https://*","http://*"},
		AllowedMethods 		: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		ALlowedHeaders 		: []string{"*"},
		ExposedHeaders 		: []string{"Link"},
		AllowCredentials 	: false,
		MaxAge				: 300,
	})

	server := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("server started on %v",portString)
	servErr := server.ListenAndServe()
	if servErr != nil {
		log.Fatal(servErr)
	}
	

}
