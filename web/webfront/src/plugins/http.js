import Vue from 'vue'
import axios from 'axios'

axios.defaults.baseURL = 'https://mzqcgc.cn/api/v1'
Vue.prototype.$http = axios
