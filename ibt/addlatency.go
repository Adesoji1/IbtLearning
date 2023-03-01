/// Add latency metric to our application code:
import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/shifts", shiftsHandler)

	go http.ListenAndServe(":8080", nil)

	// measure network latency
	for {
		resp, err := http.Get("http://localhost:8080/users")
		if err != nil {
			log.Printf("error measuring network latency: %v", err)
			continue
		}
		resp.Body.Close()
		duration := resp.Header.Get("X-Duration")
		ms, err := strconv.ParseInt(duration, 10, 64)
		if err != nil {
			log.Printf("error parsing network latency: %v", err)
			continue
		}
		latency.Set(float64(ms) / 1000.0)
		time.Sleep(10 * time.Second)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// ...
	duration := time.Since(start)
	w.Header().Set("X-Duration", strconv.FormatInt(int64(duration/time.Millisecond), 10))
	// ...
}
