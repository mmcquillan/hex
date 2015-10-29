package commands

import (
	"github.com/mmcquillan/jane/bambooapi"
)

func Build(user string, pass string, msg string) (results string) {

	key := "?"
	switch msg {
	case "client":
		key = "SYN-SCTR"
	case "cloud":
		key = "SYN-SCDR"
	case "admin":
		key = "SYN-SANR"
	case "html":
		key = "SYN-SHR"
	}
	if key != "?" {
		bambooapi.Queue("prysminc.atlassian.net", user, pass, key)
		results = "Queued Build for " + msg
	} else {
		results = "Not sure what build that is, try: [ client | cloud | admin | html ]"
	}
	return results

}
