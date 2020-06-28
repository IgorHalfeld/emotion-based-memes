import Vue from 'vue'
import VueRouter from 'vue-router'

const TakePicture = () => import('../views/TakePicture.vue')
const Loading = () => import('../views/Loading.vue')
const MemeSelector = () => import('../views/MemeSelector.vue')
const Result = () => import('../views/Result.vue')

Vue.use(VueRouter)

const routes = [
  {
    path: '/take',
    name: 'TakePicture',
    component: TakePicture
  },
  {
    path: '/processing',
    name: 'Processing',
    component: Loading
  },
  {
    path: '/select',
    name: 'MemeSelector',
    component: MemeSelector
  },
  {
    path: '/result',
    name: 'Result',
    component: Result
  },
  {
    path: '*',
    redirect: { name: 'TakePicture' }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
