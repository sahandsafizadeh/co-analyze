<template>
  <table>
    <thead>
    <tr>
      <th>#</th>
      <th>Country</th>
      <th>New Deaths</th>
      <th>Total Deaths</th>
      <th>New Cases</th>
      <th>Total Cases</th>
      <th>Doses Administered</th>
      <th>Fully Vaccinated</th>
      <th>Vaccination Percentage</th>
      <th>Population</th>
    </tr>
    </thead>
    <tbody>
    <base-statistics-table-row v-for="(s, i) in statistics"
                               :key="i"
                               :i="i + 1"
                               :country="s.country"
                               :new-deaths="s.newDeaths"
                               :total-deaths="s.totalDeaths"
                               :new-cases="s.newCases"
                               :total-cases="s.totalCases"
                               :doses-administered="s.dosesAdministered"
                               :fully-vaccinated="s.fullyVaccinated"
                               :vaccination-percentage="s.vaccinationPercentage"
                               :population="s.population"
    ></base-statistics-table-row>
    </tbody>
  </table>
</template>

<script>
import BaseStatisticsTableRow from "./BaseStatisticsTableRow";

export default {
  components: {
    BaseStatisticsTableRow,
  },

  props: {
    statistics: {
      type: Array,
      required: true,
    },
  },
};
</script>

<style scoped>
table {
  margin: 40px auto;
  border-radius: 5px;
  border: 7px ridge #69ab00;
}

th {
  position: sticky;
  top: 0;
  padding: 7px;
  border-radius: 3px;
  border: 3px solid #095484;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.5);
  text-align: center;
  font-size: 0.85rem;
  background-color: #d9d9d9;
}
</style>