// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import App from './App'
import router from './router'
import axios from 'axios'

Vue.config.productionTip = false

Vue.use(BootstrapVue);

/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    components: { App },
    template: '<App/>',
    /*data() {
      return {
      }
    },
    mounted: function () {
      axios.get(`http://localhost:8080/rest`)
      .then(response => {
        //console.log("before: " + this.server_data);
        this.server_data = response.data
        //console.log("customers: " + response.data);

        //console.log("after: " + this.server_data);
      })
      .catch(e => {
        this.errors.push(e)
      })
    },*/
})