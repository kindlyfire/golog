<template>
    <div>
        <p>{{ message }}</p>
        <router-view></router-view>
    </div>
</template>

<script>
module.exports = {
    data: function() {
        return {
            message: "Test !"
        }
    }
}
</script>