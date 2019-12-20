package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/fcgi"
	"net/http/httputil"
	"os"
	"runtime"

	"github.com/BjoernSchilberg/tonne/abfallberatungen"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var appAddr string
var signingKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	signingKey = []byte(os.Getenv("SIGNINGKEY"))
	appAddr = os.Getenv("APPADDR")
}

func requestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		//log.Println(string(requestDump))
		log.Println(fmt.Sprintf("%q", requestDump))
	})
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return signingKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {

			fmt.Fprintf(w, "Not Authorized!")
		}
	})
}

func main() {
	mux := http.NewServeMux()

	var err error

	if appAddr != "" {
		// Run as a local web server
		//mux.HandleFunc("/termine", isAuthorized termine.GetTermine(db)))
		mux.HandleFunc("/tonne/abfallberatungen", abfallberatungen.Get())
		log.Println("Listening on " + appAddr + "...")
		//err = http.ListenAndServe(appAddr, requestLogger(mux))
		// cors.Default() setup the middleware with default options being
		// all origins accepted with simple methods (GET, POST). See
		// documentation below for more options.
		handler := cors.Default().Handler(mux)
		err = http.ListenAndServe(appAddr, handler)
	} else {
		// Run as FCGI via standard I/O
		mux.HandleFunc("/fcgi-bin/tonne/abfallberatungen", abfallberatungen.Get())
		err = fcgi.Serve(nil, mux)
	}
	if err != nil {
		log.Fatal(err)
	}

}
