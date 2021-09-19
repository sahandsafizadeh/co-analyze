package repository

import (
	"server/server/src/config"
	"server/server/src/model"

	"server/server/src/util"
	"time"
)

// ErrUniqueCountryName Used to prevent duplicate country insertion.
// Based on SQLite unique constraint.
const ErrUniqueCountryName = "UNIQUE constraint failed: Countries.name"

// ErrUniqueLatestRecordForCountry Used to prevent duplicate current date-country statistics insertion.
// Based on SQLite unique constraint.
const ErrUniqueLatestRecordForCountry = "UNIQUE constraint failed: DateStatistics.country, DateStatistics.date"

// SaveGeneralStatisticsOfCountry Inserts a new country entity with its statistics into database.
// Any occurring database error is returned.
func SaveGeneralStatisticsOfCountry(stat model.Statistics) error {
	db := config.DB
	_, err := db.Exec(`INSERT INTO Countries (
                       name,
                       population,
                       totalCases,
                       totalDeaths,
                       fullyVaccinated,
                       dosesAdministered,
                       vaccinationPercentage)
                       VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		stat.Country,
		stat.Population,
		stat.TotalCases,
		stat.TotalDeaths,
		stat.FullyVaccinated,
		stat.DosesAdministered,
		stat.VaccinationPercentage)
	return err
}

// SaveLatestStatisticsOfCountry Inserts a new date-country entity for the latest statistics into database.
// Any occurring database error is returned.
func SaveLatestStatisticsOfCountry(stat model.Statistics, now time.Time) error {
	db := config.DB
	_, err := db.Exec(`INSERT INTO DateStatistics (
                       date,
                       country,
                       newCases,
                       newDeaths)
                       VALUES ($1, $2, $3, $4)`,
		now.String(),
		stat.Country,
		stat.NewCases,
		stat.NewDeaths)
	return err
}

// UpdateGeneralStatisticsForCountry Updates country entity general statistics in database.
// Any occurring database error is returned.
func UpdateGeneralStatisticsForCountry(stat model.Statistics) error {
	db := config.DB
	_, err := db.Exec(`UPDATE Countries SET
                     population = $1,
                     totalCases = $2,
                     totalDeaths = $3,
                     fullyVaccinated = $4,
                     dosesAdministered = $5,
                     vaccinationPercentage = $6
                     WHERE name = $7`,
		stat.Population,
		stat.TotalCases,
		stat.TotalDeaths,
		stat.FullyVaccinated,
		stat.DosesAdministered,
		stat.VaccinationPercentage,
		stat.Country)
	return err
}

// UpdateLatestStatisticsForCountry Updates the latest statistics of the date-country entity in database.
// Any occurring database error is returned.
func UpdateLatestStatisticsForCountry(stat model.Statistics, now time.Time) error {
	db := config.DB
	_, err := db.Exec(`UPDATE DateStatistics SET
                       newCases = $1,
                       newDeaths = $2
                       WHERE date = $3 AND country = $4`,
		stat.NewCases,
		stat.NewDeaths,
		now.String(),
		stat.Country)
	return err
}

// FindAllStatistics Selects general and current statistics for each country.
// Current date is first extracted for filtering operation.
// Floating point field is rounded.
// Any occurring database error is returned.
func FindAllStatistics() (stats []model.Statistics, err error) {
	db := config.DB
	now := util.GetTodayDate()
	rows, err := db.Query(`SELECT
       name,
       newCases,
       newDeaths,
       population,
       totalCases,
       totalDeaths,
       fullyVaccinated,
       dosesAdministered,
       vaccinationPercentage
			FROM Countries c INNER JOIN DateStatistics d
           	ON c.name = d.country
	   WHERE date = $1
	   ORDER BY vaccinationPercentage DESC`, now.String())
	if err != nil {
		return stats, err
	}
	defer rows.Close()

	allStatistics := make([]model.Statistics, 0)
	for rows.Next() {
		var stat model.Statistics
		err := rows.Scan(&stat.Country,
			&stat.NewCases,
			&stat.NewDeaths,
			&stat.Population,
			&stat.TotalCases,
			&stat.TotalDeaths,
			&stat.FullyVaccinated,
			&stat.DosesAdministered,
			&stat.VaccinationPercentage)
		if err != nil {
			return stats, err
		}
		stat.VaccinationPercentage = util.RoundTo2DecimalPlaces(stat.VaccinationPercentage)
		allStatistics = append(allStatistics, stat)
	}
	if err := rows.Err(); err != nil {
		return stats, err
	}
	return allStatistics, nil
}
