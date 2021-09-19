<template>
  <div class="bar">
    <div class="bar-content"
         :style="barContentStyle"
         @mouseout="hideInfo"
         @mouseover="showInfo"
    ></div>
    <transition>
      <base-bar-info v-if="mouseOn"
                     :x="x"
                     :y="y"
                     :color="color"
                     :statistics="statistics"
                     :vaccination="vaccination"
                     :country="country"
                     :type="type"
      ></base-bar-info>
    </transition>
  </div>
</template>

<script>
import BaseBarInfo from "./BaseBarInfo";

export default {
  components: {
    BaseBarInfo,
  },

  props: {
    height: {
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
      x: null,
      y: null,
      mouseOn: false,
    };
  },

  computed: {
    barContentStyle() {
      return {
        backgroundColor: this.color,
        height: this.height + this.UNIT,
      };
    },
  },

  methods: {
    showInfo(e) {
      this.x = e.clientX;
      this.y = e.clientY;
      this.mouseOn = true;
    },
    hideInfo() {
      this.mouseOn = false;
    },
  },
};
</script>

<style scoped>
.bar {
  height: 100%;
  margin-right: 1px;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
}

.bar-content {
  width: 5px;
  border-radius: 3px;
}

.bar-content:hover {
  transition: all 0.1s ease-out;
  transform: scaleX(1.5) translateY(-5px);
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}

.v-enter-active,
.v-leave-active {
  transition: all 50ms ease;
}

.v-enter-to,
.v-leave-from {
  opacity: 1;
}
</style>