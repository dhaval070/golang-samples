package auth;
import (
    "context"
    "net/http"
    "log"
    "gossa/handlers"
)

func WithAuth (next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        log.Println(r.RequestURI)

        if r.RequestURI == "/login" || r.RequestURI == "/auto-login" {
            next.ServeHTTP(w, r)
            return
        }

        token, err := handlers.Verify(w, r)
        if err != nil {
            http.Error(w, err.Error(), 401)
            return
        }

        ctx := context.WithValue(r.Context(), "props", token.Claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

