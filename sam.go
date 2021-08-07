package main

import (
	"github.com/gocql/gocql"
	"fmt"
	"net/http"
	"os"
	"log"
       )


func homePage(w http.ResponseWriter, r *http.Request){
        o , e := os.Hostname()
	cluster := gocql.NewCluster("cassandra")
	cluster.Consistency = gocql.Quorum
 	cluster.ProtoVersion = 4
    cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"}
 	_, err := cluster.CreateSession()
	if err == nil && e == nil {
		fmt.Fprintf(w, "cassandra_up{instance=%q} 1\n",o)
		} else {
		fmt.Fprintf(w, "cassandra_up{instance=%q} 0\n",o)
    }
}

func handleRequests() {
    http.HandleFunc("/metrics", homePage)
    log.Fatal(http.ListenAndServe(":9101", nil))
}


func main() {
          handleRequests()
    }
