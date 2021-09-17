<template>
  <div class="root">
    <div v-if="error" id="error" class="content">Error :(</div>
    <div v-else id="loader" class="content">{{ content }}</div>
  </div>
</template>

<script>
export default {
  props: {
    error: {
      type: Boolean,
      required: true,
    },
  },

  data() {
    return {
      CONTENT_CHANGE_TIMEOUT_MS: 500,
      FULL_CONTENT: '...',
      content: '',
    };
  },

  created() {
    setInterval(() => {
      if (this.content === this.FULL_CONTENT)
        this.content = '';
      else
        this.content += '.';
    }, this.CONTENT_CHANGE_TIMEOUT_MS);
  },
};
</script>

<style scoped>
.root {
  z-index: 2;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: #000000;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.content {
  text-align: center;
  font-weight: bold;
  font-size: 6rem;
}

#error {
  color: #ff0000;
}

#loader {
  color: #ffffff;
}
</style>