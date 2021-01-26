import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/login.vue'
import Admin from '../views/admin.vue'
import Index from '../components/admin/Index.vue'
import AddArt from '../components/article/addarticle.vue'
import ArtList from '../components/article/articlelist.vue'
import CaeList from '../components/category/categorylist.vue'
import UserList from '../components/user/userlist.vue'
import IndexCom from '../views/comm/Index'
import ArticleCom from '../views/comm/Article.vue'
import ArtOpenCom from '../views/comm/ArtOpen.vue'
import CategoryCom from '../views/comm/Category.vue'
import CateArt from '../views/comm/CateArt.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'IndexCom',
    component: IndexCom
  },
  {
    path: '/Article',
    name: 'ArticleCom',
    component: ArticleCom
  },
  {
    path: '/ArtOpenCom',
    name: 'ArtOpenCom',
    component: ArtOpenCom
  },
  {
    path: '/CategoryCom',
    name: 'CategoryCom',
    component: CategoryCom
  },
  {
    path: '/CateArt',
    name: 'CateArt',
    component: CateArt
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/admin',
    name: 'admin',
    component: Admin,
    children: [
      { path: 'index', component: Index },
      { path: 'addart1', component: AddArt },
      { path: 'addart/:id', component: AddArt, props: true },
      { path: 'artlist', component: ArtList },
      { path: 'catelist', component: CaeList },
      { path: 'userlist', component: UserList }
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (to.path === '/') return next()
  if (to.path === '/Article') return next()
  if (to.path === '/ArtOpenCom') return next()
  if (to.path === '/CategoryCom') return next()
  if (to.path === '/CateArt') return next()
  if (to.path === '/login') return next()
  if (!token && to.path !== '/login') {
    next('/login')
  } else {
    next()
  }
})

export default router
