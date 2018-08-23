<template>
    <div class="panel-container">
        <p>Welcome to your Golog instance !</p>
    </div>
</template>

<script>
module.exports = {
    beforeCreate: function() {
        const wm = this.$root

        wm.$emit('titleChange', "Dashboard")
    }
}
</script>

<style>

</style>
