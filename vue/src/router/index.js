import { createRouter, createWebHistory } from 'vue-router'
import Layout from '../Layout/Layout'
//路由
//要进行二次路由，嵌套路由
const routes = [
  {
    path: '/',// “/”为主页
    name: 'Layout',
    component: Layout,
    redirect:"/login",//路由的自动跳转,网址的输入
    children:[
      {
        path:'user',
        name:'User',
        component:()=>import("@/views/User")//导入位置语法
      },

      {
        path: '/book',//book为url的网址
        name: 'book',
        component:() => import("@/views/Book")
      },


    ]
  },

  {
    path: '/login',
    name: 'login',
    component:() => import("@/views/login")
  },

  {
    path: '/Register',
    name: 'Register',
    component:() => import("@/views/Register")
  },

]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
