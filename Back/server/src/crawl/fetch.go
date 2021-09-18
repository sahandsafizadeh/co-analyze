package crawl

import (
	"golang.org/x/net/html"
	"net/http"

	"server/server/src/model"
)

const VaccinationStatisticsSourceURL = "https://covidvax.live/"
const deathsAndCaseSourceURL = "https://www.worldometers.info/coronavirus/"

func Fetch() (dcStats, vacStats []model.Statistics, err error) {
	dcHtml, err := getHTML(deathsAndCaseSourceURL)
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
