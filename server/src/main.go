package main

import(
	"fmt"
	"log"
	"net/http"
	"os"
)

var active_zone string  

func zoneHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form err: %v", err)
		return
	}
	zone := r.FormValue("zone")
	//turn off the currently active zone
	file_string := fmt.Sprintf("/sys/class/gpio/gpio%s/value", active_zone)
	f, err := os.Create(file_string)
	if err != nil {
		fmt.Println(err, file_string)
		return
	}

	f.WriteString("0")
	f.Close()

	file_string = fmt.Sprintf("/sys/class/gpio/gpio%s/value", zone)
	fmt.Printf("file to open %s\n", file_string)
	f, err = os.Create(file_string)
	if err != nil {
		fmt.Println(err, file_string)
		return
	}

	f.WriteString("1")
	f.Close()
	
	active_zone = zone
	//turn on the corresponding GPIO line
	fmt.Fprintf(w, "Zone %s activated!\n", zone)
}



func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/activate_zone", zoneHandler)

	fmt.Printf("Starting irrigation server on port 65000\n")
	active_zone = "2"

	if err := http.ListenAndServe(":65000", nil); err != nil {
		log.Fatal(err)
	}
}