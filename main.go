package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

// setInterval call p function every "interval" time
func setInterval(p interface{}, interval time.Duration) chan<- bool {
	ticker := time.NewTicker(interval)
	stopIt := make(chan bool)
	go func() {

		for {

			select {
			case <-stopIt:
				fmt.Println("stop setInterval")
				return
			case <-ticker.C:
				// Run()
				reflect.ValueOf(p).Call([]reflect.Value{})
			}
		}

	}()

	// return the bool channel to use it as a stopper
	return stopIt
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	fmt.Println("Started!")
	port := os.Getenv("PORT")
	// start the setInterval and call handler every 2 seconds
	stopper := setInterval(Run, 2*time.Second)
	fmt.Println("running ...")

	http.HandleFunc("/", homePage)
	http.ListenAndServe(":"+port, nil)

	// avoid the unused variable warn:
	_ = stopper
	// to stop setInterval uncomment the next line:
	// stopper <- true

	// pause the console
	<-make(chan bool)

}

// func main() {
// 	fmt.Println("Hi, this is cmd version")
// 	for {
// 		Run()
// 	}
// 	//port := os.Getenv("PORT")
// 	//
// 	//// HandleAndSendMailCmd()
// 	//if port == "" {
// 	//	port = "8080"
// 	//}
// 	//http.HandleFunc("/", HandleAndSendMail)
// 	//http.ListenAndServe(":"+port, nil)

// }
