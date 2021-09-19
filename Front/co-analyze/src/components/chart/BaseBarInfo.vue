<template>
  <div id="container" :style="containerStyle">
    <div class="center-bold-text">{{ country }}</div>
    <hr/>
    <section class="info">
      <span id="vaccination" class="center-bold-text">{{ vaccination }}%</span>
      <span class="unit">&nbsp;Vaccinated</span>
    </section>
    <section class="info">
      <span class="center-bold-text" :style="statisticStyle">{{ statistics }}</span>
      <span class="unit">&nbsp;{{ type }}</span>
    </section>
  </div>
</template>

<script>
export default {
  props: {
    x: {
      type: Number,
      required: true,
    },
    y: {
      type: Number,
      required: true,
    },
    color: {
      type: String,
      required: true,
    },
    statistics: {
      type: Number,
      required: true,
    },
    vaccination: {
      type: Number,
      required: true,
    },
    country: {
      type: String,
      required: true,
    },
    type: {
      type: String,
      required: true,
      validator(value) {
        return value === 'Cases' || value === 'Deaths';
      },
    },
  },

  data() {
    return {
      UNIT: 'px',
      TOP_MARGIN: 25,
      LEFT_MARGIN: 7,
    };
  },

  computed: {
    containerStyle() {
      let y = this.y + this.TOP_MARGIN;
      let x = this.x + this.LEFT_MARGIN;
      return {
        top: y + this.UNIT,
        left: x + this.UNIT,
      };
    },
    statisticStyle() {
      return {
        color: this.color,
      };
    },
  },
};
</script>

<style scoped>
#container {
  z-index: 1;
  position: fixed;
  padding: 5px;
  border-radius: 3px;
  background-color: #ffffff;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.5);
}

.center-bold-text {
  text-align: center;
  font-weight: bold;
  font-size: 0.9rem;
}

.info {
  margin: 2px auto;
  text-align: center;
}

.unit {
  font-size: 0.45rem;
}

#vaccination {
  color: #007c00;
}
</style>