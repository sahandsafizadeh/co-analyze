<template>
  <input type="text"
         placeholder="Type Country"
         @input="$emit('update:modelValue', $event.target.value)"/>
</template>

<script>
export default {
  props: {
    modelValue: {
      type: String,
      required: false,
    },
  },
};
</script>

<style scoped>
input {
  display: block;
  margin: 100px auto 0;
  padding-bottom: 10px;
  outline: none;
  border: none;
  border-bottom: 3px groove;
  font-size: 1.1rem;
}

input:focus {
  border-bottom-color: #0000e5;
  font-style: normal;
  font-weight: bold;
}
</style>