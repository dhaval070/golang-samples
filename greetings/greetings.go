package greetings
import (
    "fmt"
    "errors"
    "time"
    "math/rand"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {

    if name == "" {
        return "", errors.New("name is required")
    }
    message := fmt.Sprintf(getRandomMsg(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    msgs := make(map[string]string)

    for _, name := range names {
        msg, err := Hello(name)

        if err != nil {
            return nil, err
        }

        msgs[name] = msg

    }

    return msgs, nil
}

func getRandomMsg() string {
    greets := []string {
        "Hi  %v, how r u",
        "holla %v",
        "by by %v",
    }

    return greets[rand.Intn(len(greets))]
}
