import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { Config } from 'primevue'
import Aura from '@primevue/themes/aura'
import 'primeicons/primeicons.css'

createApp(App).use(Config, {
    theme: {
        preset: Aura
    }
}).mount('#app')
