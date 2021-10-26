import {createRouter,createWebHistory,createWebHashHistory } from "vue-router"

import Redirect from '../components/Redirect.vue'
import Main from '../components/Main.vue'

const router = createRouter({
  history:createWebHistory(),
  routes:[
      {
          path:"/",
          component: Main
      },
      {
          path:"/:key",
          component: Redirect
      }
  ]
});

export default router
