import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import './plugins/http'
import moment from 'moment'
import './assets/prism.css'

Vue.config.productionTip = false

Vue.filter('dateformat', function (indate, outdate) {
  return moment(indate).format(outdate)
})

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
