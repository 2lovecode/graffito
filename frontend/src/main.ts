import { createApp } from 'vue'

import router from './router'
import './style.css'
import App from './App.vue'
import './worker'

import { createPinia } from 'pinia';
import { createHead } from '@unhead/vue'
import { registerSW } from 'virtual:pwa-register';
import { plausible } from './plugins/plausible.plugin';

// import 'virtual:uno.css';

import { naive } from './plugins/naive.plugin';
import { i18nPlugin } from './plugins/i18n.plugin';

const app = createApp(App)

app.use(router)


app.use(createPinia());
app.use(createHead());
app.use(i18nPlugin);
app.use(router);
app.use(naive);
app.use(plausible);

app.mount('#app')
