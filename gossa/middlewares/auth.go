package auth;
import (
    "net/http"
    "fmt"
    "log"
    "gossa/handlers"
)

func WithAuth (next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        log.Println(r.RequestURI)

        if r.RequestURI == "/login" {
            next.ServeHTTP(w, r)
            return
        }

        if handlers.Verify(w, r) {
            fmt.Println("valid")
            next.ServeHTTP(w, r)
        } else {
            fmt.Println("Invalid")
        }
    })
}

