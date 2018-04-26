package main

import (
	"strings"
)

func main2() {
	s := "appinsights://client:key"

	parseAppInsightMoniker(s)
}

func parseAppInsightMoniker(filename string) (string, string, bool) {
	const appInsightPrefix = "appinsights://"

	if isAppInsights := strings.HasPrefix(filename, appInsightPrefix); isAppInsights {
		parameters := strings.TrimPrefix(filename, appInsightPrefix)
		components := strings.Split(parameters, ":")

		if len(components) != 2 {
			panic("inccorect syntax for application insights")
		}

		clientName, clientKey := components[0], components[1]

		return clientName, clientKey, true
	}

	return "", "", false
}
