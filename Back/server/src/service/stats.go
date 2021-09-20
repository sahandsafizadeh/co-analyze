package service

import (
	"server/server/src/crawl"
	"server/server/src/model"
	"server/server/src/repository"

	"server/server/src/config"
	"server/server/src/util"
	"strings"
)

func updateStatistics() {
	dcStats, vacStats, err := crawl.Fetch()
	if err != nil {
		config.ServiceLogger.Println(err)
	}
	fullStats := integrateStatistics(dcStats, vacStats)
	if err := updateStatisticDatabase(fullStats); err != nil {
		config.ServiceLogger.Println(err)
	}
	config.ServiceLogger.Println("Update Successful")
}

func updateStatisticDatabase(fullStatistics []model.Statistics) error {
	now := util.GetTodayDate()
	for _, stat := range fullStatistics {
		if err := repository.SaveGeneralStatisticsOfCountry(stat); err != nil &&
			err.Error() != repository.ErrUniqueCountryName {
			return err
		}
	}
	for _, stat := range fullStatistics {
		if err := repository.UpdateGeneralStatisticsForCountry(stat); err != nil {
			return err
		}
	}
	for _, stat := range fullStatistics {
		if err := repository.SaveLatestStatisticsOfCountry(stat, now); err != nil &&
			err.Error() != repository.ErrUniqueLatestRecordForCountry {
			return err
		}
	}
	for _, stat := range fullStatistics {
		if err := repository.UpdateLatestStatisticsForCountry(stat, now); err != nil {
			return err
		}
	}
	return nil
}

// ----- helpers

func integrateStatistics(dcStats, vacStats []model.Statistics) []model.Statistics {
	fullStats := make([]model.Statistics, 0)
	nameToStructMap := makeLowercaseNameToStructMap(dcStats)
	for _, vacStat := range vacStats {
		if dcStat, ok := getDCStatistics(vacStat, nameToStructMap); !ok {
			continue
		} else {
			fullStats = append(fullStats, model.Statistics{
				Country:               vacStat.Country,
				NewCases:              dcStat.NewCases,
				NewDeaths:             dcStat.NewDeaths,
				Population:            dcStat.Population,
				TotalCases:            dcStat.TotalCases,
				TotalDeaths:           dcStat.TotalDeaths,
				FullyVaccinated:       vacStat.FullyVaccinated,
				DosesAdministered:     vacStat.DosesAdministered,
				VaccinationPercentage: vacStat.VaccinationPercentage,
			})
		}
	}
	return fullStats
}

func makeLowercaseNameToStructMap(stats []model.Statistics) map[string]model.Statistics {
	nameToStructMap := make(map[string]model.Statistics)
	for _, stat := range stats {
		nameToStructMap[strings.ToLower(stat.Country)] = stat
	}
	return nameToStructMap
}

/* countries with unregistered vaccination statistics:
   01- Bahamas
   02- Benin
   03- Bulgaria
   04- Burkina Faso
   05- Chad
   06- China
   07- Comoros
   08- Egypt
   09- Ethiopia
   10- Iraq
   11- Kuwait
   12- Lesotho
   13- Liberia
   14- Moldova
   15- Philippines
   16- Qatar
   17- Singapore
   18- Syria
   19- Tanzania
   20- Uganda
   21- Uzbekistan
   22- Zambia
*/
func getDCStatistics(stat model.Statistics, nameToStructMap map[string]model.Statistics) (dcStat model.Statistics, ok bool) {
	dcStat = nameToStructMap[strings.ToLower(stat.Country)]
	if dcStat.Country == "" {
		switch stat.Country {
		case "United Arab Emirates":
			dcStat = nameToStructMap["uae"]
		case "United Kingdom":
			dcStat = nameToStructMap["uk"]
		case "Turks and Caicos Islands":
			dcStat = nameToStructMap["turks and caicos"]
		case "United States":
			dcStat = nameToStructMap["usa"]
		case "Sint Maarten (Dutch part)":
			dcStat = nameToStructMap["sint maarten"]
		case "Curacao":
			dcStat = nameToStructMap["curaçao"]
		case "South Korea":
			dcStat = nameToStructMap["s korea"]
		case "Cape Verde":
			dcStat = nameToStructMap["cabo verde"]
		case "Timor":
			dcStat = nameToStructMap["timor-leste"]
		case "Saint Vincent and the Grenadines":
			dcStat = nameToStructMap["st vincent grenadines"]
		case "Democratic Republic of Congo":
			dcStat = nameToStructMap["drc"]
		}
	}
	if dcStat.Country != "" {
		ok = true
	}
	return dcStat, ok
}
