package commands

import ()

func Deploy(url string) (results string) {

	results = "Sad trombone - Atlassian has yet to make a deployment API, so Matt is thinking about this.\nMeanwhile, go here: https://" + url + "/builds/deploy/viewAllDeploymentProjects.action"
	return results

}
