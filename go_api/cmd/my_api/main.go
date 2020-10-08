package main

import (
	"log"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

var db *pg.DB

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Did not load .env file: %s. Using given environment variables", err.Error())
	}
	if len(os.Getenv("JWT_SIGNING_KEY")) < 32 {
		log.Fatal("Please set env var JWT_SIGNING_KEY to a random string at least 32 characters in length!")
		os.Exit(1)
	}
	pgOpts := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
	}
	db = pg.Connect(pgOpts)
	defer db.Close()

	n := initRoutes()

	envPort := os.Getenv("SERVER_PORT")
	if len(envPort) == 0 {
		envPort = "8080"
	}
	n.Run(":" + envPort)
}

func initRoutes() *negroni.Negroni {
	r := mux.NewRouter()
	ar := mux.NewRouter()

	mw := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	r.HandleFunc("/sanity", Sanity).Methods("GET")

	// How to do authenticated route
	// ar.HandleFunc("/api/friends/accept/{userID}", AcceptFriend).Methods("POST")
	an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(ar))

	if os.Getenv("DEBUG_MODE_ENABLE") == "true" {
		c := cors.AllowAll()
		an.Use(c)
		r.PathPrefix("/api").Handler(an)
		n := negroni.Classic()
		n.Use(c)
		n.UseHandler(r)
		return n
	} else {
		r.PathPrefix("/api").Handler(an)
		n := negroni.Classic()
		n.UseHandler(r)
		return n
	}
}
