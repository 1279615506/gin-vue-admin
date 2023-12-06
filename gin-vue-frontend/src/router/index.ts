import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import UserLogin from '../pages/user/UserLogin.vue'
import UserInfo from '../pages/user/UserInfo.vue'

const routes: Array<RouteRecordRaw> = [
  {path: '/',component: HomeView,
   children:[
    
  ]
  },
  {path:'/userinfo', component: UserInfo},
  {path:'/login', component: UserLogin},
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
