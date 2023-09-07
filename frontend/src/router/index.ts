import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../components/HomePage.vue'
import SandBox from '../components/SandBox.vue'


const routes = [
    { path: '/', component:  HomePage},
    { path: '/sandbox', component: SandBox },
  ]
export default createRouter({
    history: createWebHistory(),
    routes: routes
})