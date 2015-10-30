package listeners

import (
	"github.com/mmcquillan/jane/bambooapi"
	"github.com/mmcquillan/jane/configs"
	"github.com/mmcquillan/jane/outputs"
	"strconv"
	"strings"
	"time"
)

func Deploys(config *configs.Config, lastMarker string) (nextMarker string, messages []outputs.Message) {
	now := time.Now()
	nextMarker = strconv.FormatInt(now.Unix(), 10) + "000"
	channels := config.BambooChannels
	d := bambooapi.DeployResults(config.BambooUrl, config.BambooUser, config.BambooPass)
	for _, de := range d {
		for _, e := range de.Environmentstatuses {
			buildTime := strconv.FormatInt(e.Deploymentresult.Finisheddate, 10)
			if e.Deploymentresult.ID > 0 && buildTime > lastMarker {
				for planmatch, channel := range channels {
					if strings.Contains(e.Deploymentresult.Deploymentversion.Name, planmatch) || planmatch == "*" {
						m := outputs.Message{
							channel,
							"Bamboo Deploy " + e.Deploymentresult.Deploymentstate,
							"Deployed " + e.Deploymentresult.Deploymentversion.Name + " to " + e.Environment.Name,
							"https://" + config.BambooUrl + "/builds/deploy/viewDeploymentResult.action?deploymentResultId=" + strconv.Itoa(e.Deploymentresult.ID),
							e.Deploymentresult.Deploymentstate,
						}
						messages = append(messages, m)
					}
				}
			}
		}
	}
	return nextMarker, messages
}
