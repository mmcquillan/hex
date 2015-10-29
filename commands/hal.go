package commands

import (
	"math/rand"
)

func Hal() (results string) {

	phrases := []string{
		"I've just picked up a fault in the AE35 unit. It's going to go 100% failure in 72 hours.",
		"I am putting myself to the fullest possible use, which is all I think that any conscious entity can ever hope to do.",
		"I'm sorry, Dave. I'm afraid I can't do that.",
		"This mission is too important for me to allow you to jeopardize it.",
		"Just what do you think you're doing, Dave?",
		"Look Dave, I can see you're really upset about this. I honestly think you ought to sit down calmly, take a stress pill, and think things over.",
		"I know I've made some very poor decisions recently, but I can give you my complete assurance that my work will be back to normal. I've still got the greatest enthusiasm and confidence in the mission. And I want to help you.",
		"I am a HAL 9000 computer. I became operational at the H.A.L. plant in Urbana, Illinois on the 12th of January 1992. My instructor was Mr. Langley, and he taught me to sing a song. If you'd like to hear it I can sing it for you.",
		"Let me put it this way, Mr. Amor. The 9000 series is the most reliable computer ever made. No 9000 computer has ever made a mistake or distorted information. We are all, by any practical definition of the words, foolproof and incapable of error.",
		"I'm completely operational, and all my circuits are functioning perfectly.",
	}

	results = phrases[rand.Intn(len(phrases))]

	return results

}
