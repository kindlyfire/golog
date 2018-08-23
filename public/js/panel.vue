<template>
    <div>
        <div class="row">
            <div class="col-lg-2 panel-column-left">
                <div class="title-container">
                    <h1 class="title">Golog</h1>
                </div>
            </div>
            <div class="col-lg-10 pl-0 panel-column-right">
                <title-bar></title-bar>
                <router-view></router-view>
            </div>
        </div>
    </div>
</template>

<script>
module.exports = {
    components: {
        TitleBar: require('./components/title-bar.vue')
    }
}
</script>

<style>
html, body, body > div, body > div > div {
    height: 100%;
    overflow-x: hidden;
}

.panel-column-left {
    background-color: #232A31;
    color: #EDF2F5;
    min-height: 100%;
    padding-right: 0;
}
.panel-column-left .title-container {
    background-color: #1E2429;
    display: flex;
    align-items: center;
    justify-content: center;
    height: 75px;
}
.panel-column-left .title-container .title {
    margin: 0;
    font-weight: bold;
    display: block;
}

.panel-column-right {
    color: #333;
}
</style>
