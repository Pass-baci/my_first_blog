import Vue from 'vue'
import VueRouter from 'vue-router'
const Login = () => import(/* webpackChunkName: "Login" */ '../views/login.vue')
const Admin = () => import(/* webpackChunkName: "Admin" */ '../views/admin.vue')

// 页面路由组件
const Index = () => import(/* webpackChunkName: "Index" */ '../components/admin/Index.vue')
const AddArt = () => import(/* webpackChunkName: "addarticle" */ '../components/article/addarticle.vue')
const ArtList = () => import(/* webpackChunkName: "ArtList" */ '../components/article/articlelist.vue')
const CateList = () => import(/* webpackChunkName: "CateList" */ '../components/category/categorylist.vue')
const UserList = () => import(/* webpackChunkName: "UserList" */ '../components/user/userlist.vue')

// 路由重复点击捕获错误
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push (location, onResolve, onReject) {
  if (onResolve || onReject) return originalPush.call(this, location, onResolve, onReject)
  return originalPush.call(this, location).catch(err => err)
}

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    meta: {
      title: '请登录'
    },
    component: Login
  },
  {
    path: '/',
    name: 'admin',
    meta: {
      title: '后台管理页面'
    },
    component: Admin,
    children: [
      {
        path: 'index',
        component: Index,
        meta: {
          title: '后台管理页面'
        }
      },
      {
        path: 'addart1',
        component: AddArt,
        meta: {
          title: '新增文章'
        }
      },
      {
        path: 'addart/:id',
        component: AddArt,
        meta: {
          title: '编辑文章'
        },
        props: true
      },
      {
        path: 'artlist',
        component: ArtList,
        meta: {
          title: '文章列表'
        }
      },
      {
        path: 'catelist',
        component: CateList,
        meta: {
          title: '分类列表'
        }
      },
      {
        path: 'userlist',
        component: UserList,
        meta: {
          title: '用户列表'
        }
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()

  const userToken = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!userToken) {
    next('/login')
  } else {
    next()
  }
})

export default router
