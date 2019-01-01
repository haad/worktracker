<template>
  <div class="customers">
    <h3>List of defined customers</h3>

    <b-table striped hover :items=server_data :fields=fields :bordered=true :sort-by.sync="Rate"></b-table>
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
                {key: 'Rate', sortable: true},
                {key: 'ContactName', sortable: false},
                {key: 'ContactEmail', sortable: false}
              ]
    }
  },

  created() {
    axios.get(`http://localhost:8080/rest`)
    .then(response => {
      this.server_data = response.data
    })
    .catch(e => {
      this.errors.push(e)
    })
  }
}
</script>