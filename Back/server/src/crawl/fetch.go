package crawl

import (
	"golang.org/x/net/html"
	"net/http"

	"server/server/src/model"
)

const VaccinationStatisticsSourceURL = "https://covidvax.live/"
const DeathsAndCaseStatisticsSourceURL = "https://www.worldometers.info/coronavirus/"

// Fetch Gathers Covid-19 statistics as HTML content and parses them.
// Vaccination statistics are fetched from {VaccinationStatisticsSourceURL}.
// Deaths and cases statistics are fetched from {DeathsAndCaseStatisticsSourceURL}.
// Then the required information is extracted and returned for the two HTML documents.
// Any error due to retrieving HTML contents is returned.
func Fetch() (dcStats, vacStats []model.Statistics, err error) {
	dcHtml, err := getHTML(DeathsAndCaseStatisticsSourceURL)
	if err != nil {
		return dcStats, vacStats, err
	}
	vacHtml, err := getHTML(VaccinationStatisticsSourceURL)
	if err != nil {
		return dcStats, vacStats, err
	}
	dcStats = parseDeathsAndCasesStatistics(dcHtml)
	vacStats = parseVaccinationStatistics(vacHtml)
	return dcStats, vacStats, nil
}

// ----- helpers

func getHTML(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	htmlBody, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}
	return htmlBody, nil
}
