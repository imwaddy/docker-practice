<template>
  <div id="app">
    <h1>{{ greetingMessage }}</h1>
    
    <label for="inputNumber">Enter a number:</label>
    <input type="number" id="inputNumber" v-model="userInput" @input="handleInput">

    <p>Entered Number: {{ userInput }}</p>
    <p>Is Even? {{ isEven }}</p>
    <p>Squared: {{ squareValue }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      greetingMessage: 'Hello Vue.js!',
      userInput: 0,
    };
  },
  computed: {
    isEven() {
      return this.userInput % 2 === 0;
    },
    squareValue() {
      return this.userInput ** 2;
    }
  },
  methods: {
    handleInput() {
      // Additional operations can be performed when the input changes
      console.log('Input changed:', this.userInput);
    }
  }
};
</script>

<style>
@import './style.css'

/* Additional styles specific to App.vue can be added here */
</style>
