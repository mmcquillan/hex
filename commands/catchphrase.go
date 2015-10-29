package commands

import (
	"math/rand"
)

func CatchPhrase() (results string) {

	phrases := []string{
		"I'll be back.",
		"Stop moving my cheese.",
		"Embrace the change.",
		"We are all six shell scripts away from losing our job.",
		"Automate all the things.",
		"Frankly, my dear, I don't give a damn,",
		"I'm going to make him an offer he can't refuse.",
		"I coulda been a contender.",
		"I've got a feeling we're not in Kansas anymore.",
		"Go ahead, make my day.",
		"May the Force be with you.",
		"I love the smell of napalm in the morning.",
	}

	results = phrases[rand.Intn(len(phrases))]

	return results

}
