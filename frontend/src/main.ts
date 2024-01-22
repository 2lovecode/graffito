import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import AppRouter from './router'
import AppPinia from './store'



const app = createApp(App)

app.use(ElementPlus)
app.use(AppRouter)
app.use(AppPinia)

app.mount('#app')
