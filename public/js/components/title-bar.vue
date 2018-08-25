<template>
    <div>
        <div class="title-bar-container">
            <h2 class="title">
                {{ title }}
            </h2>
            <button type="button" class="btn btn-primary" v-if="button" @click="button.callback">
                {{ button.text }}
            </button>
        </div>
        <div class="title-bar-error" v-for="(err, i) in errors" :key="i">
            {{ err }}
            <span @click="close_err(i)">(close)</span>
        </div>
    </div>
</template>

<script>
module.exports = {
    data() {
        const wm = this.$root

        return {
            title: "Dashboard",
            loading: false,
            button: null,
            errors: []
        }
    },
    created() {
        const wm = this.$root

        wm.$on('titlebar_change', (params) => {
            for (let p of Object.keys(params)) {
                if (['title', 'loading', 'button'].indexOf(p) !== -1) {
                    this[p] = params[p]
                }
            } 
        })
        wm.$on('show_error', (msg) => {
            this.errors.push(msg)
        })
    },
    methods: {
        close_err(id) {
            this.errors.pop(id)
        }
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

.title-bar-error {
    background-color: #D32F2F;
    color: white;
    padding: 10px;
    padding-left: 20px;
}
.title-bar-error span {
    padding-left: 10px;
    cursor: pointer;
}
</style>
