package main

import (
	"log"
	"main/matcher"
	"math/rand"
	"time"
)

var Matcher *matcher.TradeMatcher

func init() {
	Matcher = matcher.NewMatcher()
}

func main() {

	// generate random seed global
	rand.Seed(time.Now().UTC().UnixNano())

	if Matcher != nil {
		err := Matcher.Start("tcp", "0.0.0.0:8000")
		if err != nil {
			log.Println(err)
		}
	}

	finish := make(chan bool)
	<-finish
}
