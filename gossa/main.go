package main
import (
    "github.com/joho/godotenv"
    "gossa/handlers"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    auth "gossa/middlewares"
)

func init() {
    godotenv.Load()
}

func main() {
    router := mux.NewRouter()
    router.Use(auth.WithAuth)

    router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "hello\n")
    })
    router.HandleFunc("/auto-login", handlers.AutoLogin)
    router.HandleFunc("/login", handlers.Login).Methods("POST")
    router.HandleFunc("/locations", handlers.GetLocations).Methods("GET")
    router.HandleFunc("/locations/leagues/{id}", handlers.LocLeagues).Methods("GET")
    router.HandleFunc("/locations/re-assign", handlers.ReAssign).Methods("POST")
    //id = 0 to insert record
    router.HandleFunc("/locations/{id}", handlers.EditLocation).Methods("GET", "POST")

    fmt.Println("starting api")
    log.Fatal(http.ListenAndServe(":5000", router))
}

