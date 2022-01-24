import Vue from 'vue'
import VueRouter from 'vue-router'
import 'element-ui/lib/theme-chalk/index.css'
import ElementUI from 'element-ui'

import App from './App.vue'
import Login from "@/components/Login";
import Register from "@/components/Register";
import Chat from "@/components/Chat";
import Welcome from "@/components/Welcome";

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: '欢迎',
    component: Welcome,
    meta:{
      title: '欢迎'
    }
  },
  {
    path: '/Login',
    name: '登录',
    component: Login,
    meta: {
      title: '登录'
    }
  },
  {
    path: '/Register',
    name: '注册',
    component: Register,
    meta: {
      title: '注册'
    }
  },
  {
    path: '/Chat',
    name: '聊天',
    component: Chat,
    meta: {
      title: '聊天'
    }
  }
]

const router = new VueRouter(
    {
      routes
    }
);

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title;
  }
  next()
})

new Vue({
  router: router,
  render: h => h(App),
}).$mount('#app')
