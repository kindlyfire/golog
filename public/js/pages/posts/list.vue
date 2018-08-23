<template>
    <div class="panel-container">
        <p>Post listing !</p>
    </div>
</template>

<script>
module.exports = {
    beforeCreate: function() {
        const wm = this.$root

        wm.$emit('titleChange', "Post list")
    }
}
</script>

<style>

</style>
