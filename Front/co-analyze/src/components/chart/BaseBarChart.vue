<template>
  <section class="bar-chart">
    <h3 id="y-label" class="legend">New {{ type }}</h3>
    <base-bar v-for="(s, i) in statistics"
              :key="i"
              :height="heights[i]"
              :color="color"
              :statistics="getStat(s)"
              :vaccination="s.vaccinationPercentage"
              :country="s.country"
              :type="type"
    ></base-bar>
    <h3 id="x-label" class="legend">
      &#8594;&#8594;&#8594;&#8594;&#8594;
      Vaccination
      &#8594;&#8594;&#8594;&#8594;&#8594;
    </h3>
    <div id="max-stat" class="legend quantity">{{ maxStat }}</div>
    <div id="min-perc" class="legend quantity">{{ minPercentage }}%</div>
    <div id="max-perc" class="legend quantity">{{ maxPercentage }}%</div>
  </section>
</template>

<script>
import BaseBar from "./BaseBar";

export default {
  components: {
    BaseBar,
  },

  props: {
    type: {
      type: String,
      required: true,
      validator(value) {
        return value === 'Cases' || value === 'Deaths';
      },
    },
    stats: {
      type: Array,
      required: true,
      validator(value) {
        return value.length <= 255;
      },
    },
  },

  data() {
    return {
      MIN_HEIGHT: 10,
      MAX_HEIGHT: 500,
      CASE_COLOR: '#ffdd00',
      DEATH_COLOR: '#ff0000',
      minPercentage: 0,
      maxPercentage: 0,
      maxStat: 0,
      heights: [],
      statistics: [],
    };
  },

  watch: {
    stats() {
      this.reinitializeChart();
    },
  },

  computed: {
    color() {
      return this.type === 'Cases' ? this.CASE_COLOR : this.DEATH_COLOR;
    },
  },

  methods: {
    getStat(fullStatus) {
      return this.type === 'Cases' ? fullStatus.newCases : fullStatus.newDeaths;
    },
    initMaxStat() {
      let max = this.getStat(this.stats[0]);
      for (let i = 1; i < this.stats.length; i++) {
        let cur = this.getStat(this.stats[i]);
        if (max < cur)
          max = cur;
      }
      this.maxStat = 0 < max ? max : 1;
    },
    initStatSortedOnVaccination() {
      for (let stat of this.stats)
        this.statistics.push(stat);
      this.statistics.sort((a, b) => a.vaccinationPercentage - b.vaccinationPercentage);
    },
    initScaledHeights() {
      for (let stat of this.statistics) {
        let rawHeight = (this.MAX_HEIGHT * this.getStat(stat)) / this.maxStat;
        let height = this.MIN_HEIGHT + Math.floor(rawHeight);
        this.heights.push(height);
      }
    },
    reinitializeChart() {
      if (this.stats.length === 0)
        return;
      this.initMaxStat();
      this.initStatSortedOnVaccination();
      this.initScaledHeights();
      this.minPercentage = this.statistics[0].vaccinationPercentage;
      this.maxPercentage = this.statistics[this.statistics.length - 1].vaccinationPercentage;
    },
  },

  created() {
    this.reinitializeChart();
  },
};
</script>

<style scoped>
.bar-chart {
  position: relative;
  height: 520px;
  width: fit-content;
  padding: 0 1px;
  margin: 105px auto;
  border-radius: 5px;
  border: solid gray 2px;
  background-color: #000000;
  display: flex;
  flex-direction: row;
  justify-content: center;
}

.legend {
  position: absolute;
}

#y-label {
  top: -47px;
}

#x-label {
  bottom: -53px;
}

.quantity {
  color: #757575;
  font-weight: bold;
}

#max-stat {
  left: 0;
  top: -25px;
}

#min-perc {
  left: 0;
  bottom: -25px;
}

#max-perc {
  right: 0;
  bottom: -25px;
}
</style>