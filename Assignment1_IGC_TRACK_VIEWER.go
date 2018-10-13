

// assignment 1 golang cloud tech

package main 


import "fmt"
import "net/http"
//import "encoding/json"
import "strings"	
import "time"
import "log"

//import "github.com/marni/goigc"
//import "github.com/gorilla/mux"

type IGCTrack struct {

	ID int 					`json:"track_id"`
	Pilot string 			`json:"pilot"`
	Glider string  			`json:"glider"`
	Glider_id string `	     json:"glider_id"`
	Track_length float64	`json:"track_length"`
 	H_date time.Time 		`json:"h_date"`

}



type IGCTrack struct {
	HDate       time.Time `json:"H_date"`
	Pilot       string    `json:"pilot"`
	Glider      string    `json:"glider"`
	GliderID    string    `json:"glider_id"`
	ID          string    `json:"track_id"`
	TrackLength float64   `json:"track_length"`
	Data        igc.Track `json:"-"`
}


// gets start time
var start = time.Now() 


/*
// TODO psydo stuff
type URL struct {	
	url string
}

db map[string]IGC


*/



/*

//* from marni code
func returnAllTracks(w http.ResponseWriter, db StudentStorage) {
if db.Count() == 0 {
		json.NewEncoder(w).Encode([]Student{})
	} else {
		a := make([]Student, 0, db.Count())
		for _, s := range db.GetAll() {
			a = append(a, s)
		}
		json.NewEncoder(w).Encode(a)
}
}
*/
//













func handler1(w http.ResponseWriter, r *http.Request ) {

switch r.Method	{	


case "POST": {
	if r.Body == nil {
		http.Error(w, "IGC post request needs to have a JSON body ", http.StatusBadRequest)
		return
	}


/*
	var i IGC 
	err := json.NewDecoder(r.Body).Decode(&i) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO
	// need to also process the url igc file by using marni's code?..
	//



	// check if the url / igc-file is new
_, ok := db.Get(i.glider_id)
	if ok{
		http.Error(w, "This glider_id already exitst", http.StatusBadRequest)
		return
	}

// new igc, add them to local storage 
//	db.add(i)
	fmt.Fprint(w, i.glider_id)

	return

*/
}

case "GET": {


//http.Header.Add(w.Header(), "content-type", "application/json")



parts := strings.Split(r.URL.Path, "/")

/*
// checks if the given url is valid 
//
// http://localhost:8080/igcinfo/               -> 404
// http://localhost:8080/<rubbish>              -> 404
// http://localhost:8080/igcinfo/api/<rubbish>  -> 404
//

*/

fmt.Println(len(parts))
fmt.Println(r.URL.Path)



// Error handling 
//
if parts[1] != "igcinfo" {
	
	http.Error(w, "Try adding igcinfo to the url \n", http.StatusBadRequest)
	return
}

if parts[2] != "api"  {

 	http.Error(w, "Hint: add /api/ :) \n", http.StatusBadRequest)
	return
}

if parts[3] == "igc" && len(parts) > 3 {
	// should return all tracks 
	// for i : tr
	return
}



	
elapsed := time.Since(start)


fmt.Fprintf(w,"uptime: %s\n", elapsed )
fmt.Fprint(w,"Info: Service for IGC tracks.\n")
fmt.Fprint(w,"version: v1")
return 
	}

default: {

	http.Error(w, "Only post and get requests are currently accepted\n", http.StatusBadRequest)
		return
	}

}


}



/*
func handlerPostNewIGC(w http.ResponseWriter, r *http.Request ) {




}*/


func main() {

//var start = time.Now()


router := mux.NewRouter()
	
//http.HandleFunc("/igcinfo/api/igc", handlerPostNewIGC)

http.HandleFunc("/", handler1)


log.Fatal(http.ListenAndServe(":8080", router))




	
}