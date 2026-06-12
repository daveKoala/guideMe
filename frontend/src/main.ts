import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import Aura from '@primeuix/themes/aura'
import 'primeicons/primeicons.css'
import './styles/base.css'
import { registerSW } from 'virtual:pwa-register'

import App from './App.vue'
import router from './router/index.ts'

// Register the service worker (autoUpdate): SW updates only swap cached
// JS/CSS/HTML — localStorage (people/trips/docs) is untouched, so auto-reload is safe.
registerSW({ immediate: true })

// Background-cache demo doc images without blocking first load. The Workbox
// CacheFirst runtime rule stores whatever the SW sees these fetches return.
function warmDocCache() {
  for (const url of [
    '/test-docs/passport.png',
    '/test-docs/boardingpass.png',
    '/test-docs/medicalcard.png',
  ]) {
    fetch(url).catch(() => {})
  }
}
if ('requestIdleCallback' in window) requestIdleCallback(() => warmDocCache())
else setTimeout(warmDocCache, 3000)

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(PrimeVue, {
  theme: {
    preset: Aura,
    options: {
      darkModeSelector: false,
    },
  },
})

app.mount('#app')
