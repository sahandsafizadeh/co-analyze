/* SQLite Database */
/* stats.sqlite */


DROP TABLE IF EXISTS Countries;
DROP TABLE IF EXISTS DateStatistics;


CREATE TABLE Countries
(
    name                  TEXT PRIMARY KEY NOT NULL,
    population            INTEGER          NOT NULL,
    totalCases            INTEGER          NOT NULL,
    totalDeaths           INTEGER          NOT NULL,
    fullyVaccinated       INTEGER          NOT NULL,
    dosesAdministered     INTEGER          NOT NULL,
    vaccinationPercentage REAL             NOT NULL,
    CONSTRAINT check_country CHECK
        (
            (LENGTH(name) <= 50)
            AND
            (0 < population)
            AND
            (0 <= totalCases)
            AND
            (0 <= totalDeaths)
            AND
            (0 <= fullyVaccinated)
            AND
            (0 <= dosesAdministered)
            AND
            (vaccinationPercentage BETWEEN 0 AND 100)
        )
);

CREATE TABLE DateStatistics
(
    date      TEXT    NOT NULL,
    country   INTEGER NOT NULL,
    newCases  INTEGER NOT NULL,
    newDeaths INTEGER NOT NULL,
    PRIMARY KEY (country, date),
    FOREIGN KEY (country)
        REFERENCES Countries (name)
        ON DELETE CASCADE,
    CONSTRAINT check_statistic CHECK
        (
            (length(date) <= 100)
            AND
            (0 <= newCases)
            AND
            (0 <= newDeaths)
        )
);
