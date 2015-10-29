package commands

import (
	"math/rand"
)

func Jane() (results string) {

	phrases := []string{
		"\"But I have eyes,\" she said. \"And ears. I see everything in all the Hundred Worlds. I watch the sky through a thousand telescopes. I overhear a trillion conversations every day.\" She giggled a little. \"I'm the best gossip in the universe.\"",
		"\"Twisted and perverse are the ways of the human mind,\" Jane intoned. \"Pinocchio was such a dolt to try to become a real boy. He was much better off with a wooden head.\"",
		"\"It's the most charming thing about humans. You are all so sure that the lesser animals are bleeding with envy because they didn't have the good fortune to be born Homo sapiens.\"",
	}

	results = phrases[rand.Intn(len(phrases))]

	return results

}
