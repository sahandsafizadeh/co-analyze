package model

type Statistics struct {
	Country               string  `json:"country"`
	NewCases              int64   `json:"newCases"`
	NewDeaths             int64   `json:"newDeaths"`
	Population            int64   `json:"population"`
	TotalCases            int64   `json:"totalCases"`
	TotalDeaths           int64   `json:"totalDeaths"`
	FullyVaccinated       int64   `json:"fullyVaccinated"`
	DosesAdministered     int64   `json:"dosesAdministered"`
	VaccinationPercentage float64 `json:"vaccinationPercentage"`
}
