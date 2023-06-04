package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello you have requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
}

func main() {
	http.HandleFunc("/", rootHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("The web server has started")
	http.ListenAndServe(":8080", nil)
}

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