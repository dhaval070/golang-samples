package handlers
import (
    "log"
    "fmt"
    "gossa/models"
    "encoding/json"
    "net/http"
    "io"
    "strings"
    "github.com/gorilla/mux"
    "strconv"
    "gossa/db"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "os"
)

func GetLocations(w http.ResponseWriter, r *http.Request) {
    var result []models.Location

    db := db.Db()

    res, err := db.Query(`select id, location, status, copy_method, ifnull(locked_by_league,'') locked_by_league,
        ifnull(locked_by_event_id, 0) locked_by_event_id, use_rclone from location order by location`)

    if err != nil {
        http.Error(w, err.Error(), 500)
    }
    for res.Next() {
        var loc models.Location

        err := res.Scan(&loc.Id, &loc.Location, &loc.Status, &loc.CopyMethod, &loc.LockedByLeague,
            &loc.LockedByEventId, &loc.UseRclone)

        if (err != nil) {
            log.Fatal(err)
        }

        result = append(result, loc)
    }
    json.NewEncoder(w).Encode(result)
}

func ReAssign(w http.ResponseWriter, r *http.Request) {
    db := db.Db()

    buf, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("errorx: %v", err)
        http.Error(w, err.Error(), 500)
    }

    var body struct {
        Id int `json:"id"`
        NewId int `json:"new_id"`
        Leagues []string `json:"leagues"`
    }

    if err = json.Unmarshal(buf, &body); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    for _, l := range body.Leagues {
        var s strings.Builder
        fmt.Fprintf(&s, `update gos_%s.location set global_id=? where global_id=?`, l)
        res, err := db.Query(s.String(), body.NewId, body.Id)

        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        fmt.Fprintf(w, "success %v", res)
    }
}

func EditLocation(w http.ResponseWriter, r *http.Request) {
    db := db.Db()

    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
    var loc models.Location

    w.Header().Set("content-type", "application/json")

    if r.Method == "GET" {
        res, err := db.Query(`select id, location, status, copy_method,
        ifnull(locked_by_league,'') locked_by_league, ifnull(locked_by_event_id, 0) locked_by_event_id,
        use_rclone from location where id = ?`, id)

        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }

        if res.Next() {
            err := res.Scan(&loc.Id, &loc.Location, &loc.Status, &loc.CopyMethod, &loc.LockedByLeague,
                &loc.LockedByEventId, &loc.UseRclone)

            if (err != nil) {
                log.Fatal(err)
            }
        }
        json.NewEncoder(w).Encode(loc)
        return
    }

    buf, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("error: %v", err)
        http.Error(w, err.Error(), 500)
    }

    var body models.Location

    if err = json.Unmarshal(buf, &body); err != nil {
        log.Println("err in body parse")
        http.Error(w, err.Error(), 500)
        return
    }
    // Edit location
    if id != 0 {
        _, err := db.Query(`update location set location=?, status=?, copy_method=?,use_rclone=?
            where id=?`, body.Location, body.Status, body.CopyMethod, body.UseRclone, id)

        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        fmt.Fprint(w, `{"success":1}`)
        return
    }
    // Add location
    _, err = db.Query(`insert into location(location, status, copy_method, use_rclone) values(
        ?, ?, ?, ?)`, body.Location, body.Status, body.CopyMethod, body.UseRclone)

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    json.NewEncoder(w).Encode(body)
}

func Login(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")
    db := db.Db()

    var body struct {
        Email string `json:"email"`
        Password string `json:"password"`
        OpsDb string `json:"opsDb"`
    }

    buf, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("error: %v", err)
        http.Error(w, err.Error(), 500)
    }

    if err = json.Unmarshal(buf, &body); err != nil {
        log.Println("err in body parse")
        http.Error(w, err.Error(), 500)
        return
    }

    //body.Email = "dhaval070@gmail.com"
    //body.Password = ""
    //body.OpsDb = "gos"

    res, err := db.Query(`select id, email, password from user where email=? and superuser=1`, body.Email)

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    var id int
    var email,password string

    if (!res.Next()) {
        http.Error(w, "User not found", 401)
        return
    }
    if err := res.Scan(&id, &email, &password); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(body.Password)); err != nil {
        http.Error(w, err.Error(), 401)
    }

    type Payload struct {
        Id int `json:"id"`
        Email string `json:"email"`
        OpsDb string `json:"opsDb"`
    }
    var data = Payload {
        id,
        email,
        body.OpsDb,
    }

    token := jwtToken(id, email, body.OpsDb)
    c := http.Cookie {
        Name: "token",
        Path: "/",
        Value: token,
        HttpOnly: true,
    }

    fmt.Println(token)

    w.Header().Set("Set-Cookie", c.String())
    json.NewEncoder(w).Encode(data)

}

func jwtToken(uid int, email string, opsDb string) string {
    claim := jwt.MapClaims{ "id": uid, "email": email, "opsDb": opsDb }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

    var SECRET = os.Getenv("JWT_SECRET")
    tokenStr, err := token.SignedString([]byte(SECRET))
    if err != nil {
        fmt.Println(err.Error())
        return ""
    }
    return tokenStr
}

func Verify (w http.ResponseWriter, r *http.Request) bool {
    c, err := r.Cookie("token")

    if err != nil {
        http.Error(w, err.Error(), 401)
        return false
    }

    var token *jwt.Token
    var SECRET = os.Getenv("JWT_SECRET")

    token, err = jwt.Parse(c.Value, func (token *jwt.Token) (interface {}, error) {
        return []byte(SECRET), nil
    })

    if err != nil {
        http.Error(w, err.Error(), 500)
        return false
    }

    return token.Valid
}

func AutoLogin(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")
    db := db.Db()

    var body struct {
        Email string `json:"email"`
        Password string `json:"password"`
        OpsDb string `json:"opsDb"`
    }

    body.Email = "dhaval070@gmail.com"
    body.Password = ""
    body.OpsDb = "gos"

    res, err := db.Query(`select id, email from user where email=? and superuser=1`, body.Email)

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    var id int
    var email string

    if (!res.Next()) {
        http.Error(w, "User not found", 401)
        return
    }
    if err := res.Scan(&id, &email); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    type Payload struct {
        Id int `json:"id"`
        Email string `json:"email"`
        OpsDb string `json:"opsDb"`
    }
    var data = Payload {
        id,
        email,
        body.OpsDb,
    }

    token := jwtToken(id, email, body.OpsDb)
    c := http.Cookie {
        Name: "token",
        Path: "/",
        Value: token,
        HttpOnly: true,
    }

    w.Header().Set("Set-Cookie", c.String())
    json.NewEncoder(w).Encode(data)

}

