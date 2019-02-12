<template>
  <div class="customers">
    <div class="second_nav">
      <b-navbar toggleable="md" type="light" variant="light">
        <b-navbar-toggle target="nav_collapse"></b-navbar-toggle>
        <b-navbar-brand href="/app/customers">Customer projects:</b-navbar-brand>

        <b-collapse is-nav id="nav_collapse">
          <b-navbar-nav>
            <b-nav-item v-for="cust in server_data"><router-link :to="{ name: 'Project', params: { id: cust.ID }}">{{ cust.Name }}</router-link></b-nav-item>
          </b-navbar-nav>
        </b-collapse>
      </b-navbar>
    </div>
    <div class=customers_data>
      <h3>List of defined customers</h3>

      <b-table striped hover :items=server_data :fields=fields :bordered=true>
        <template slot="Projects" slot-scope="server_data">
          <router-link :to="{ name: 'Project', params: { id: server_data.item.ID }}">Projects</router-link>
        </template>
      </b-table>
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
                {key: 'Rate', sortable: true},
                {key: 'ContactName', sortable: false},
                {key: 'ContactEmail', sortable: false},
                {key: 'Projects', label: 'Projects' }
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