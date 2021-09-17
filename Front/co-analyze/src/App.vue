<template>
  <app-loader v-if="loading" :error="facedError"></app-loader>
  <base-bar-chart type="Deaths" :stats="chartStats"></base-bar-chart>
  <base-bar-chart type="Cases" :stats="chartStats"></base-bar-chart>
  <hr/>
  <base-text-input v-model="country"></base-text-input>
  <app-statistics-table :statistics="tableStats"></app-statistics-table>
</template>

<script>
import axios from "axios";
import AppLoader from "./components/AppLoader";
import BaseBarChart from "./components/chart/BaseBarChart";
import BaseTextInput from "./components/input/BaseTextInput";
import AppStatisticsTable from "./components/table/AppStatisticsTable";

export default {
  components: {
    AppLoader,
    BaseBarChart,
    BaseTextInput,
    AppStatisticsTable,
  },

  data() {
    return {
      HOST_URL: 'http://localhost:8080/',
      chartStats: [],
      tableStats: [],
      statistics: [],
      country: null,
      loading: true,
      facedError: false,
    };
  },

  watch: {
    country(value) {
      this.tableStats = this.statistics
          .filter(s => s.country.startsWith(value));
    },
  },

  methods: {
    load() {
      axios.get(this.HOST_URL)
          .then((response) => {
            this.statistics = response.data;
            this.chartStats = this.statistics;
            this.tableStats = this.statistics;
            this.loading = false;
          })
          .catch(() => {
            this.facedError = true;
          });
    },
    loadRandomly() {
      let countries = [
        'Australia',
        'Canada',
        'China',
        'France',
        'Iran',
        'Iraq',
        'Italy',
        'Japan',
        'Russia',
        'South Africa',
        'UAE',
        'United Kingdom',
        'USA',
        'Vietnam',
      ];
      setTimeout(() => {
        for (let i = 0; i < 220; i++)
          this.statistics.push({
            country: countries[Math.floor(Math.random() * countries.length)],
            newDeaths: Math.floor(Math.random() * 1000),
            newCases: Math.floor(Math.random() * 20000),
            totalDeaths: Math.floor(Math.random() * 200000),
            totalCases: Math.floor(Math.random() * 5000000),
            population: Math.floor(Math.random() * 2000000000),
            vaccinationPercentage: Math.floor(Math.random() * 100),
            fullyVaccinated: Math.floor(Math.random() * 100000000),
            dosesAdministered: Math.floor(Math.random() * 2000000000),
          });
        this.chartStats = this.statistics;
        this.tableStats = this.statistics;
        this.loading = false;
      }, 2500);
    },
  },

  created() {
    // this.load();
    this.loadRandomly();
  },
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Azeret+Mono:wght@300&display=swap');

* {
  font-family: 'Azeret Mono', serif;
}
</style>
