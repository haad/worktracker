<template>
  <div id="app">

    <!-- Fixed navbar -->
    <nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
      <a class="navbar-brand" href="#">Fixed navbar</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul class="nav navbar-nav">
          <router-link tag="li" active-class="active" to="/" exact><a>Home</a></router-link>
          <router-link tag="li" active-class="active" to="/about"><a>About</a></router-link>
          <router-link tag="li" active-class="active" to="/rest"><a>Rest</a></router-link>
        </ul>
      </div>
    </nav>

    <router-link :to="{ name: 'Hello' }">Home</router-link>
    <router-link to="/about">About</router-link>
    <router-link to="/rest">Rest</router-link>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 60px;
  margin-left: 100px;
}

body {
  margin: 0;
}

main {
  text-align: center;
  margin-top: 40px;
}

header {
  margin: 0;
  height: 56px;
  padding: 0 16px 0 24px;
  background-color: #35495E;
  color: #ffffff;
}

header span {
  display: block;
  position: relative;
  font-size: 20px;
  line-height: 1;
  letter-spacing: .02em;
  font-weight: 400;
  box-sizing: border-box;
  padding-top: 16px;
}

.customers {
  column-count: 1;
  column-gap: 1em;
  margin-top: 100px;
}

</style>
