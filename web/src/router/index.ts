import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../components/HomePage.vue'
import ToolsMain from '../components/tools/ToolsMain.vue'
import Sandbox from '../components/tools/Sandbox.vue'

// 2. 定义一些路由
// 每个路由都需要映射到一个组件。
// 我们后面再讨论嵌套路由。
const routes = [
  {
    path: '/',
    component: HomePage 
  },
  { 
    path: '/tools',
    children:[
      {
        path: '',
        component: ToolsMain,
      },
      {
        path: 'sandbox',
        component: Sandbox,
      }
    ]
  },
]

// 3. 创建路由实例并传递 `routes` 配置
// 你可以在这里输入更多的配置，但我们在这里
// 暂时保持简单
const AppRouter = createRouter({
  history: createWebHistory(),
  routes, // `routes: routes` 的缩写
})

export default AppRouter