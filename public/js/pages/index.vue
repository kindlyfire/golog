<template>
    <div class="panel-container">
        <p>Welcome to your Golog instance !</p>
    </div>
</template>

<script>
import TitleBar from '@/components/title-bar.vue'

module.exports = {
    beforeCreate: function() {
        const wm = this.$root

        wm.$emit('titlebar_change', {
            title: "Dashboard"
        })
    },
    components: {
        TitleBar
    }
}
</script>

<style>

</style>
