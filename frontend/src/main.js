/**
 * main.js
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components

import { createPinia } from 'pinia'
// Composables
import { createApp, markRaw } from 'vue'
import { registerPlugins } from '@/plugins'

import App from './App.vue'
// Plugins
import vuetify from './plugins/vuetify'
import router from './router'

const pinia = createPinia()
pinia.use(({ store }) => {
  store.$router = markRaw(router)
})

const app = createApp(App)

registerPlugins(app)

app.use(pinia).use(vuetify).use(router).mount('#app')
