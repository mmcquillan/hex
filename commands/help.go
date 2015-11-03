package commands

import (
	"github.com/mmcquillan/jane/models"
	"sort"
	"strings"
)

func Help(config *models.Config) (results string) {
	h := map[string]bool{}
	var help []string
	for _, c := range config.Commands {
		help = append(help, c.Match)
	}
	for _, v := range help {
		if _, seen := h[v]; !seen {
			help[len(h)] = v
			h[v] = true
		}
	}
	help = help[:len(h)]
	sort.Strings(help)
	results = "*Say things like:*\n"
	for _, r := range help {
		results += r + ", "
	}
	return strings.TrimRight(results, ", ")
}
