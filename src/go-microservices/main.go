package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	details "github.com/Samvit20/go-microservices/details"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")
	response := map[string]string{
		"status": "UP",
		"time":   time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running!")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := details.GetHostname()
	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIP()

	response := map[string]string{
		"hostname": hostname,
		"IP":   IP.String(),
	}
	json.NewEncoder(w).Encode(response)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/details", detailsHandler)

	log.Println("Webserver started on 8080!")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello you have requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
// }

// func main() {
// 	http.HandleFunc("/", rootHandler)

// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	log.Println("The web server has started")
// 	http.ListenAndServe(":8080", nil)
// }

// package main

// import (
// 	"fmt"
// 	"unsafe"

// 	"github.com/Samvit20/go-microservices/geometry"
// 	"rsc.io/quote"
// )

// func rectProps(length, breadth float64) (area, perimeter float64) {
// 	area = length * breadth
// 	perimeter = (length + breadth) * 2
// 	return
// }

// func main() {
// 	x := 23
// 	name := "Samvit"
// 	isWorking := false
// 	fmt.Println("Hello world!")
// 	fmt.Println(quote.Go())
// 	fmt.Println(x, name, isWorking)
// 	fmt.Printf("Type of name: %T and size: %d\n", name, unsafe.Sizeof(name))

// 	area, perimeter := rectProps(1, 2)
// 	fmt.Printf("Area: %f and Perimeter: %f\n", area, perimeter)

// 	var daysOfTheMonth = map[string]int{"Jan": 31, "Feb": 28}
// 	fmt.Println(daysOfTheMonth)

// 	area_from_geometry_package := geometry.Area(1, 2)
// 	diagonal_from_geometry_package := geometry.Diagonal(1, 2)
// 	fmt.Printf("Area: %f and DIagonal: %f\n", area_from_geometry_package, diagonal_from_geometry_package)
// }
