package main

import(
	"fmt"
	"log"
	"net/http"
)


func zoneHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form err: %v", err)
		return
	}
	zone := r.FormValue("zone")
	//turn on the corresponding GPIO line
	fmt.Fprintf(w, "Zone %s activate!\n", zone)
}



func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/activate_zone", zoneHandler)

	fmt.Printf("Starting irrigation server on port 65000\n")

	if err := http.ListenAndServe(":65000", nil); err != nil {
		log.Fatal(err)
	}
}