
require('babel-polyfill')
const Vue = require('vue')
const Router = require('vue-router')
const comp = require('./panel.vue')

const routes = [
    { path: '/', component: require('./pages/index.vue') },
    { path: '/posts', component: require('./pages/posts/list.vue') },
    { path: '/posts/create', component: require('./pages/posts/edit.vue') },
    { path: '/posts/edit/:postId', component: require('./pages/posts/edit.vue') },
]

const router = new Router({
    mode: 'history',
    base: window.PANEL_URL,
    routes
})

Vue.use(Router)

const wm = new Vue({
    router,
    el: '#app',
    render: (h) => h(comp)
})

router.beforeEach((to, from, next) => {
    wm.$emit('titleBarChange', {
        loading: false,
        button: false
    })
    next()
})

