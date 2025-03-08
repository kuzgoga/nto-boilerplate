import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { Config } from 'primevue'
import Aura from '@primevue/themes/aura'
import 'primeicons/primeicons.css'
import { ru } from 'primelocale/js/ru.js'

createApp(App).use(Config, {
    theme: {
        preset: Aura,
    },
    locale: ru
}).mount('#app')
