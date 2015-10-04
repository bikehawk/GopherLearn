package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!doctype html>
		 <html>
		 	<head>
				<title>AJAX POST</title>
				<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
			</head>
			<body>
				<div id="result">Click the button to post to the service.</div>
				<input id="postbtn" type="button" value="POST to service">
				<script>
					$(document).ready(function() {
						$("#postbtn").click(function() {
							$.ajax({
								url: 'http://localhost:8080/post',
								type: 'POST',
								contentType: 'application/json',
								dataType: 'json',
								data: JSON.stringify({'key1': 'value1', 'key2': 'value2'}),
								success: function(data) {
									$('#result').html("key1: "+data.key1+", key2: "+data.key2);
								},
							});
						});
					});
				</script>
			</body>
		</html>`
	w.Write([]byte(fmt.Sprintf(html)))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data Data
	fmt.Printf("%#v\n", r.Body)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(b))
	json.Unmarshal(b, &data)
	fmt.Printf("data: %#v\n", data)
	o, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, string(o))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", mainHandler).Methods("GET")
	router.HandleFunc("/post", postHandler).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
