<template>
  <tr>
    <td class="row-num">{{ i }}</td>
    <td class="country">{{ country }}</td>
    <td :class="newDeathsClass">{{ presentNumber(theNewDeaths) }}</td>
    <td class="side-info">{{ presentNumber(totalDeaths) }}</td>
    <td :class="newCasesClass">{{ presentNumber(theNewCases) }}</td>
    <td class="side-info">{{ presentNumber(totalCases) }}</td>
    <td class="side-info">{{ presentNumber(theDosesAdministered) }}</td>
    <td :class="fullyVaccinatedClass">{{ presentNumber(theFullyVaccinated) }}</td>
    <td :class="vaccinationPercentageClass">{{ vaccinationPercentage }}%</td>
    <td class="population">{{ presentNumber(population) }}</td>
  </tr>
</template>

<script>
export default {
  props: {
    i: {
      type: Number,
      required: true,
    },
    country: {
      type: String,
      required: true,
    },
    newDeaths: {
      type: Number,
      required: true,
    },
    totalDeaths: {
      type: Number,
      required: true,
    },
    newCases: {
      type: Number,
      required: true,
    },
    totalCases: {
      type: Number,
      required: true,
    },
    dosesAdministered: {
      type: Number,
      required: true,
    },
    fullyVaccinated: {
      type: Number,
      required: true,
    },
    vaccinationPercentage: {
      type: Number,
      required: true,
    },
    population: {
      type: Number,
      required: true,
    },
  },

  computed: {
    hasNewCases() {
      return 0 < this.newCases;
    },
    hasNewDeaths() {
      return 0 < this.newDeaths;
    },
    hasFullyVaccinated() {
      return 0 < this.fullyVaccinated;
    },
    hasDosesAdministered() {
      return 0 < this.dosesAdministered;
    },
    hasVaccinatedPercentage() {
      return 0 < this.vaccinationPercentage;
    },
    theNewCases() {
      return this.hasNewCases ? this.newCases : null;
    },
    theNewDeaths() {
      return this.hasNewDeaths ? this.newDeaths : null;
    },
    theFullyVaccinated() {
      return this.hasFullyVaccinated ? this.fullyVaccinated : null;
    },
    theDosesAdministered() {
      return this.hasDosesAdministered ? this.dosesAdministered : null;
    },
    newCasesClass() {
      return {
        'important': true,
        'new-cases': this.hasNewCases,
      };
    },
    newDeathsClass() {
      return {
        'important': true,
        'new-deaths': this.hasNewDeaths,
      };
    },
    fullyVaccinatedClass() {
      return {
        'fully-vaccinated': this.hasFullyVaccinated,
      };
    },
    vaccinationPercentageClass() {
      return {
        'important': true,
        'vaccination-percentage': this.hasVaccinatedPercentage,
      };
    },
  },

  methods: {
    presentNumber(num) {
      if (num === null)
        return null;
      if (num < 1000)
        return num;
      let presented = '';
      let str = num.toString();
      let q = str.length / 3;
      let r = str.length % 3;

      if (0 < r)
        presented = str.substring(0, r) + ',';
      while (0 <= --q) {
        let s = 3 * q;
        presented += str.substring(str.length - s - 3, str.length - s) + ',';
      }
      return presented.substring(0, presented.length - 1);
    },
  },
};
</script>

<style scoped>
td {
  padding: 5px;
  border: 1px groove;
  text-align: center;
  font-weight: bold;
}

.important {
  font-size: 1.1rem;
}

.side-info {
  font-size: 0.9rem;
}

.row-num {
  font-size: 0.8rem;
  color: #868686;
}

.country {
  color: #004579;
}

.new-deaths {
  color: #ffffff;
  background-color: #ff0000;
}

.new-cases {
  background-color: #ffff00;
}

.fully-vaccinated {
  background-color: #93ff8f;
}

.vaccination-percentage {
  color: #ffffff;
  background-color: #008100;
}

.population {
  color: #6b6b6b;
}
</style>