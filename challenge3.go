package main

import "fmt"

func main() {

	// run task 6
	task6()
}

// TASK 6: calculate the number of tweets per each unique user

// define tweet structure
type Tweet struct {
	Username string
	Message string
}

// create collection of tweets
var tweets = []Tweet{
	Tweet{"Baky", "My first tweet!"},
	Tweet{"Baky", "Yellow lives matters!"},
	Tweet{ "Simon", "After 10 years of experience, i am still surprised by some of the features of PowerBI"},
	Tweet{"Karel", "I created a new cryptocurrency"},
	Tweet{ "Simon", "I think, we should strive to support every gender-neutral movements raising in the IT domain!"},
}

func task6() {

	// create empty map
	tweetCounts := make(map[string]int)

	// TODO: use range loop, iterate over tweets
	// TODO: fill the map with number of tweets per user
	// TODO: "username" should be used as a unique key in map
	// TODO: use int to store number of tweets per user

	// print tweetCounts
	for user, numberOfTweets := range tweetCounts {
		fmt.Printf("user: %v, number of tweets: %v \n", user, numberOfTweets)
	}
}