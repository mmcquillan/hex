package connectors

import (
	"github.com/SlyMarbo/rss"
	"github.com/kennygrant/sanitize"
	"github.com/mmcquillan/bambooapi"
	"github.com/projectjane/jane/models"
	"html"
	"log"
	"strconv"
	"strings"
	"time"
)

type Bamboo struct {
}

func (x Bamboo) Listen(commandMsgs chan<- models.Message, connector models.Connector) {
	defer Recovery(connector)
	now := time.Now()
	buildMarker := ""
	deployMarker := strconv.FormatInt(now.Unix(), 10) + "000"
	for {
		buildMarker = listenBuilds(buildMarker, commandMsgs, connector)
		deployMarker = listenDeploys(deployMarker, commandMsgs, connector)
		time.Sleep(120 * time.Second)
	}
}

func (x Bamboo) Command(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	if strings.HasPrefix(message.In.Text, "bamboo status") {
		commandDeployStatus(message, publishMsgs, connector)
		commandBuildStatus(message, publishMsgs, connector)
	}
	if strings.HasPrefix(message.In.Text, "bamboo build") {
		commandBuild(message, publishMsgs, connector)
	}
}

func (x Bamboo) Publish(connector models.Connector, message models.Message, target string) {
	return
}

func (x Bamboo) Help(connector models.Connector) (help string) {
	help += "bamboo build <build key>\n"
	help += "bamboo status <environment or build key>\n"
	return help
}

func listenDeploys(lastMarker string, commandMsgs chan<- models.Message, connector models.Connector) (nextMarker string) {
	now := time.Now()
	nextMarker = strconv.FormatInt(now.Unix(), 10) + "000"
	if connector.Debug {
		log.Print("Calling deploy results with last marker: " + lastMarker)
	}
	d := bambooapi.DeployResults(connector.Server, connector.Login, connector.Pass)
	for _, de := range d {
		for _, e := range de.Environmentstatuses {
			buildTime := strconv.FormatInt(e.Deploymentresult.Finisheddate, 10)
			if e.Deploymentresult.ID > 0 && buildTime > lastMarker {
				var m models.Message
				m.Routes = connector.Routes
				m.In.Process = false
				m.Out.Text = "Bamboo Deploy " + e.Deploymentresult.Deploymentversion.Name + " to " + e.Environment.Name + " " + e.Deploymentresult.Deploymentstate
				m.Out.Detail = html.UnescapeString(sanitize.HTML(e.Deploymentresult.Reasonsummary))
				m.Out.Link = "https://" + connector.Server + "/builds/deploy/viewDeploymentResult.action?deploymentResultId=" + strconv.Itoa(e.Deploymentresult.ID)
				if e.Deploymentresult.Deploymentstate == "SUCCESS" {
					m.Out.Status = "SUCCESS"
				} else {
					m.Out.Status = "FAIL"
				}
				commandMsgs <- m
			}
		}
	}
	return nextMarker
}

func commandBuild(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	tokens := strings.Split(message.In.Text, " ")
	queue := bambooapi.Queue(connector.Server, connector.Login, connector.Pass, tokens[2])
	log.Printf("%+v", queue)
	if queue.StatusCode > 0 {
		message.Out.Text = "Problem queueing build: " + tokens[2]
	} else {
		message.Out.Text = "Queued build for " + queue.Plankey
		message.Out.Detail = "Build #" + strconv.Itoa(queue.Buildnumber)
		message.Out.Link = "https://" + connector.Server + "/builds/browse/" + queue.Plankey + "-" + strconv.Itoa(queue.Buildnumber)
	}
	publishMsgs <- message
}

func commandDeployStatus(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	tokens := strings.Split(message.In.Text, " ")
	d := bambooapi.DeployResults(connector.Server, connector.Login, connector.Pass)
	for _, de := range d {
		for _, e := range de.Environmentstatuses {
			detail := e.Deploymentresult.Deploymentversion.Name + " to " + e.Environment.Name
			if strings.Contains(strings.ToLower(detail), strings.ToLower(tokens[2])) {
				message.Out.Text = "Bamboo Deploy " + detail + " " + e.Deploymentresult.Deploymentstate
				message.Out.Detail = html.UnescapeString(sanitize.HTML(e.Deploymentresult.Reasonsummary))
				message.Out.Link = "https://" + connector.Server + "/builds/deploy/viewDeploymentResult.action?deploymentResultId=" + strconv.Itoa(e.Deploymentresult.ID)
				if e.Deploymentresult.Deploymentstate == "SUCCESS" {
					message.Out.Status = "SUCCESS"
				} else {
					message.Out.Status = "FAIL"
				}
				publishMsgs <- message
			}
		}
	}
}

func commandBuildStatus(message models.Message, publishMsgs chan<- models.Message, connector models.Connector) {
	tokens := strings.Split(message.In.Text, " ")
	b := bambooapi.BuildResults(connector.Server, connector.Login, connector.Pass)
	for _, be := range b.Results.Result {
		if strings.Contains(strings.ToLower(be.Plan.Shortname), strings.ToLower(tokens[2])) {
			message.Out.Text = "Bamboo Build " + be.Plan.Shortname
			message.Out.Detail = be.Plan.Name + " #" + strconv.Itoa(be.Buildnumber)
			message.Out.Link = be.Link.Href
			if be.Buildstate == "Successful" {
				message.Out.Status = "SUCCESS"
			} else {
				message.Out.Status = "FAIL"
			}
			publishMsgs <- message
		}
	}
}

func listenBuilds(lastMarker string, commandMsgs chan<- models.Message, connector models.Connector) (nextMarker string) {
	var displayOnStart = 0
	url := "https://" + connector.Login + ":" + connector.Pass + "@"
	url += connector.Server + "/builds/plugins/servlet/streams?local=true"
	feed, err := rss.Fetch(url)
	if err != nil {
		log.Print(err)
		return
	}
	for i := len(feed.Items) - 1; i >= 0; i-- {
		if connector.Debug {
			log.Print("Bamboo " + connector.ID + " item #" + strconv.Itoa(i) + " marker " + feed.Items[i].Date.String())
		}
		if lastMarker == "" {
			lastMarker = feed.Items[displayOnStart].Date.String()
		}
		item := feed.Items[i]
		if item.Date.String() > lastMarker {
			status := "NONE"
			if strings.Contains(item.Title, "successful") {
				status = "SUCCESS"
			}
			if strings.Contains(item.Title, "fail") {
				status = "FAIL"
			}
			var m models.Message
			m.Routes = connector.Routes
			m.In.Process = false
			m.Out.Text = "Bamboo Build " + html.UnescapeString(sanitize.HTML(item.Title))
			m.Out.Detail = html.UnescapeString(sanitize.HTML(item.Content))
			m.Out.Link = item.Link
			m.Out.Status = status
			commandMsgs <- m
			if i == 0 {
				lastMarker = item.Date.String()
			}
		}
	}
	nextMarker = lastMarker
	return nextMarker
}
