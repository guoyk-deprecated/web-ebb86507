import Vue from 'vue'
import VueRouter from 'vue-router'
import BootstrapVue from 'bootstrap-vue'
import App from './App.vue'

import './vendor/bootswatch/litera.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.config.productionTip = false

Vue.use(BootstrapVue)
Vue.use(VueRouter)

import 'vue-awesome/icons/lock'
import 'vue-awesome/icons/envelope'
import 'vue-awesome/icons/play-circle'
import 'vue-awesome/icons/brands/weibo'
import 'vue-awesome/icons/brands/twitter'
import 'vue-awesome/icons/brands/instagram'
import 'vue-awesome/icons/brands/github'
import 'vue-awesome/icons/brands/steam-square'
import Icon from 'vue-awesome/components/Icon'

Vue.component('v-icon', Icon)

import router from './router'

new Vue({
  el: '#app',
  router,
  components: {App},
  template: '<App/>'
}).$mount('#app')
