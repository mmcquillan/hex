package commands

import (
	"github.com/mmcquillan/bambooapi"
)

func Env(url string, user string, pass string) (results string) {
	results = "*Current deploys...*\n"
	d := bambooapi.DeployResults(url, user, pass)
	for _, de := range d {
		for _, e := range de.Environmentstatuses {
			if e.Deploymentresult.ID > 0 {
				results += e.Deploymentresult.Deploymentversion.Name + " deployed to " + e.Environment.Name + " - " + e.Deploymentresult.Deploymentstate + "\n"
			}
		}
	}
	return results
}
