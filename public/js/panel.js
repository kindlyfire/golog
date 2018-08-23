
const Vue = require('vue')
const Router = require('vue-router')
const comp = require('./panel.vue')

const routes = [
    { path: '/', component: require('./pages/index.vue') }
]

const router = new Router({
    mode: 'history',
    base: window.BASE_URL,
    routes
})

Vue.use(Router)

new Vue({
    router,
    el: "#app",
    render: (h) => h(comp)
})