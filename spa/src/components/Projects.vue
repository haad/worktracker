<template>
  <div class="projects" v-if="$route.params.id">
    <div class="project" v-for="cust in server_data">
      <h3 v-if="cust.ID == $route.params.id"> Projects for customer: {{ cust.Name }} </h3>
      <b-table striped hover :items=cust.Projects :fields=fields :bordered=true
               v-if="cust.ID == $route.params.id"></b-table>
    </div>
  </div>

  <div class="projects" v-else>
    <div class="project" v-for="cust in server_data">
      <h3> Projects for customer: {{ cust.Name }} </h3>
      <b-table striped hover :items=cust.Projects :fields=fields :bordered=true></b-table>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Projects',
  data() {
    return {
      server_data: [],
      errors: [],
      fields: [ {key: 'id', sortable: true},
                {key: 'name', sortable: false},
                {key: 'finished', sortable: false},
                {key: 'estimate', sortable: true},
                {key: 'logged_time', sortable: true},

              ]
    }
  },

  // Fetches posts when the component is created.
  created() {
    axios.get(`http://localhost:8080/rest`)
    .then(response => {
      // JSON responses are automatically parsed.
      this.server_data = response.data
    })
    .catch(e => {
      this.errors.push(e)
    })
  }
}
</script>