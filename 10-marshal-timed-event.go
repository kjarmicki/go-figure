package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TimedEvent struct {
	ID   int
	Time time.Time
}

func MarshalTimedEvent() {
	event := TimedEvent{
		ID:   1,
		Time: time.Date(2023, 01, 01, 01, 01, 01, 01, time.Now().Location()),
	}

	payload, _ := json.Marshal(event)
	fmt.Println((string(payload)))
}

/*
 * Default JSON marshalling of a struct can be overriden with MarshalJSON() method.
 * Embedding brings in all the properties of the struct, including its methods. Time has custom MarshalJSON() implementation.
 */
