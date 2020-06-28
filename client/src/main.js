import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// Tailwind
import './tailwind/tailwind.css'

Object.defineProperty(Vue.prototype, '$store', {
  get: () => store,
  set () {
    throw new Error('Psiu! don\'t set $store property')
  }
})

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
