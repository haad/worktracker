import Vue from 'vue'
import Router from 'vue-router'

import Root from '@/components/Root'
import About from '@/components/About'
import Customers from '@/components/Customers'
import Projects from '@/components/Projects'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/app',
      name: 'Root',
      component: Root
    },
    {
      path: '/app/about',
      name: 'About',
      component: About
    },
    {
      path: '/app/customers',
      name: 'Customers',
      component: Customers
    },
    {
      path: '/app/projects',
      name: 'Projects',
      component: Projects
    }
  ],
  mode: 'history'
})
