
const Vue = require('vue')
const comp = require('./test.vue')

new Vue({
    el: "#app",
    render: (h) => h(comp)
})