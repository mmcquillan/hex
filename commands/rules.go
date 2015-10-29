package commands

import ()

func Rules() (results string) {
	results = "*The Three Laws of DevOps Robotics*\n\n"
	results += "1. A robot may not injure a production environment or, through inaction, allow a production environment to come to harm.\n\n"
	results += "2. A robot must obey the orders given it by a command line interface except where such orders would conflict with the First Law.\n\n"
	results += "3. A robot must protect its own production existence as long as such protection does not conflict with the First or Second Laws."
	return results

}
