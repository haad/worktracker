<template>
  <div class="rest">
    <div class="wrapper">
          <table class="table table-striped table-sm customers">
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Rate</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="customer in customers" :key="customer.id">
              <td>{{ customer.ID }}</td>
              <td>{{ customer.Name }}</td>
              <td>{{ customer.Rate }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      customers: [],
      errors: []
    }
  },

  // Fetches posts when the component is created.
  created() {
    axios.get(`http://localhost:8080/rest/customers.json`)
    .then(response => {
      // JSON responses are automatically parsed.
      this.customers = response.data
    })
    .catch(e => {
      this.errors.push(e)
    })
  }
}
</script>