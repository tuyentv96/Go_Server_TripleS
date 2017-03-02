package lib

import (
	"log"
	"gopkg.in/mgo.v2"
	"sync"
)

const (
	MongoDBHosts = "127.0.0.1"
	AuthDatabase = "goinggo"
	TestDatabase = "quanlynha"
)

type Devices struct {
	Id string
	//ID	bson.ObjectId
	Dname      string
	Address     string
	Status    int
	Power	int
}

func RunQuery(query int, waitGroup *sync.WaitGroup, mongoSession *mgo.Session) {
	// Decrement the wait group count so the program knows this
	// has been completed once the goroutine exits.
	defer waitGroup.Done()

	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(TestDatabase).C("devices")

	log.Printf("RunQuery : %d : Executing\n", query)

	// Retrieve the list of stations.
	var buoyStations []Devices
	err := collection.Find(nil).All(&buoyStations)
	if err != nil {
		log.Printf("RunQuery : ERROR : %s\n", err)
		return
	}

	log.Printf("RunQuery : %d : Count[%d]\n", query, len(buoyStations))
}