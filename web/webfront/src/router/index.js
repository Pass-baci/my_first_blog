import Vue from 'vue'
import VueRouter from 'vue-router'
// 使用路由懒加载
const Home = () => import('../views/Home.vue')
const ArticleList = () => import('../components/ArticleList.vue')
const Detail = () => import('../components/Details.vue')
const Category = () => import('../components/CateList.vue')
const Search = () => import('../components/Search.vue')

Vue.use(VueRouter)

// 获取原型对象上的push函数
const originalPush = VueRouter.prototype.push
// 修改原型对象中的push方法
VueRouter.prototype.push = function push (location) {
  return originalPush.call(this, location).catch(err => err)
}

const routes = [
  {
    path: '/',
    component: Home,
    meta: { title: '欢迎来到Paradise' },
    children: [
      { path: '/', component: ArticleList, meta: { title: '欢迎来到Paradise' } },
      {
        path: 'detail/:id',
        component: Detail,
        meta: { title: '文章详情' },
        props: true
      },
      {
        path: 'category/:cid',
        component: Category,
        meta: { title: '分类信息' },
        props: true
      },
      {
        path: 'search/:title',
        component: Search,
        meta: { title: '搜索结果' },
        props: true
      }
    ]
  }
]

const router = new VueRouter({
  // mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title ? to.meta.title : '加载中'
  }
  next()
})

export default router
