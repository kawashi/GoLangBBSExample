<template>
  <div id="app">
    <h1>Go API server example.</h1>
    <BBS/>
  </div>
</template>

<script>
import BBS from './components/BBS'

export default {
  name: 'App',
  components: {
    BBS
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
