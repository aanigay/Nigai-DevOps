import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import TheChat from "@/components/chat/TheChat.vue";
import TheLogin from "@/components/login/TheLogin.vue";
import TheFeed from "@/components/feed/TheFeed.vue";
import TheDashboard from "@/components/dashboard/TheDashboard.vue";
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/login',
      name: 'login',
      component: TheLogin
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/chat',
      name: 'chat',
      component: TheChat
    },
    {
      path: '/feed',
      name: 'feed',
      component: TheFeed
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: TheDashboard
    }
  ]
})

export default router
