package crawl

import (
	"golang.org/x/net/html"
	"server/server/src/model"
)

const vaccinationTargetTableId = "countryTable"
const deathsAndCasesTargetTableId = "main_table_countries_today"

func parseVaccinationStatistics(root *html.Node) []model.Statistics {
	stats := make([]model.Statistics, 0)
	tbody := getElementById(root, vaccinationTargetTableId)
	lastRow := tbody.LastChild.PrevSibling
	for tr := tbody.FirstChild.NextSibling; ; tr = nextNode(tr, 1) {
		var stat model.Statistics

		// country
		td := tr.FirstChild.NextSibling
		if countryName, ok := extractCountryName(td); ok {
			stat.Country = countryName
		} else {
			continue
		}

		// doses administered
		td = nextNode(td, 1)
		if dosesAdministered, ok := extractInteger(td, false); ok {
			stat.DosesAdministered = dosesAdministered
		} else {
			continue
		}

		// fully vaccinated
		td = nextNode(td, 3)
		if fullyVaccinated, ok := extractInteger(td, false); ok {
			stat.FullyVaccinated = fullyVaccinated
		} else {
			continue
		}

		// vaccination percentage
		td = nextNode(td, 1)
		if vaccinationPercentage, ok := extractPercentage(td); ok {
			stat.VaccinationPercentage = vaccinationPercentage
		} else {
			continue
		}

		stats = append(stats, stat)
		if tr == lastRow {
			break
		}
	}
	return stats
}

func parseDeathsAndCasesStatistics(root *html.Node) []model.Statistics {
	stats := make([]model.Statistics, 0)
	table := getElementById(root, deathsAndCasesTargetTableId)
	tbody := nextNode(table.FirstChild.NextSibling, 1)
	lastRow := tbody.LastChild.PrevSibling
	for tr := nextNode(tbody.FirstChild.NextSibling, 8); ; tr = nextNode(tr, 1) {
		var stat model.Statistics

		// country
		td := nextNode(tr.FirstChild.NextSibling, 1)
		if countryName, ok := extractCountryName(td); ok {
			stat.Country = countryName
		} else {
			continue
		}

		// total cases
		td = nextNode(td, 1)
		if totalCases, ok := extractInteger(td, true); ok {
			stat.TotalCases = totalCases
		} else {
			continue
		}

		// new cases
		td = nextNode(td, 1)
		if newCases, ok := extractInteger(td, true); ok {
			stat.NewCases = newCases
		} else {
			continue
		}

		// total deaths
		td = nextNode(td, 1)
		if totalDeaths, ok := extractInteger(td, true); ok {
			stat.TotalDeaths = totalDeaths
		} else {
			continue
		}

		// new deaths
		td = nextNode(td, 1)
		if newDeaths, ok := extractInteger(td, true); ok {
			stat.NewDeaths = newDeaths
		} else {
			continue
		}

		// population
		td = nextNode(td, 9)
		if population, ok := extractInteger(td, false); ok {
			stat.Population = population
		} else {
			continue
		}

		stats = append(stats, stat)
		if tr == lastRow {
			break
		}
	}
	return stats
}
