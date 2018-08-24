<template>
    <div class="title-bar-container">
        <h2 class="title">
            {{ title }}
        </h2>
        <button type="button" class="btn btn-primary" v-if="button" @click="button.callback">
            {{ button.text }}
        </button>
    </div>
</template>

<script>
module.exports = {
    data: function() {
        const wm = this.$root

        return {
            title: "Dashboard",
            loading: false,
            button: null
        }
    },
    beforeCreate: function() {
        const wm = this.$root

        wm.$on('titleBarChange', (params) => {
            for (let p of Object.keys(params)) {
                if (['title', 'loading', 'button'].indexOf(p) !== -1) {
                    this[p] = params[p]
                }
            } 
        })
    }
}
</script>

<style>
.title-bar-container {
    height: 75px;
    background-color: #EDF2F5;
    position: relative;
    display: flex;
    align-items: center;
}
.title-bar-container .title {
    font-size: 35px;
    margin: 0;
    padding-left: 20px;
}
.title-bar-container button {
    margin-left: 20px;
    background-color: transparent;
    color: inherit;
}
</style>
