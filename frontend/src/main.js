import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import VueSocketIO from 'vue-3-socket.io'
import { generateUsername } from "unique-username-generator";

import App from './App.vue'
import router from './router'
import store from './store/index'
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

import "@mdi/font/css/materialdesignicons.css"

const vuetify = createVuetify({
    components,
    directives,
})

const app = createApp(App)

localStorage.removeItem('posts')
localStorage.setItem('name', generateUsername())
localStorage.setItem('role', 'guest')
localStorage.setItem('access_token', 'null')
localStorage.setItem('posts', JSON.stringify([]))

app.use(createPinia())
app.use(router)
app.use(vuetify)
app.use(store)
app.use(new VueSocketIO({
    debug: false,
    connection: 'http://localhost:8010',
    vuex: {
        store,
        actionPrefix: 'SOCKET_',
        mutationPrefix: 'SOCKET_'
    },
}))

app.mount('#app')
