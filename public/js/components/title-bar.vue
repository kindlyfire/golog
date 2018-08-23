<template>
    <div class="title-bar-container">
        <h2 class="title">
            {{ title }}
        </h2>
    </div>
</template>

<script>
module.exports = {
    data: function() {
        const wm = this.$root

        return {
            title: wm.$data.page.title
        }
    },
    beforeCreate: function() {
        const wm = this.$root

        wm.$on('titleChange', (title) => {
            this.title = title
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
</style>
