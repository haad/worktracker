<template>
  <div class="projects">
    <div class="project" v-for="cust in server_data">
      <h3> Projects for customer: {{ cust.Name }} </h3>
      <b-table striped hover :items=cust.Projects :fields=fields :bordered=true :sort-by.sync="Rate"></b-table>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      server_data: [],
      errors: [],
      fields: [ {key: 'ID', sortable: true},
                {key: 'Name', sortable: false},
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