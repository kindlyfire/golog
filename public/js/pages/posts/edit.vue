<template>
    <div>
        Editing {{ post.Title }}
    </div>
</template>

<script>
import API from '@/api'

export default {
    data() {
        return {
            postId: 0,
            post: {}
        }
    },
    created() {
        this.postId = this.$route.params.postId

        this.$root.$emit('titlebar_change', {
            title: this.postId === undefined ? 'Create post' : 'Edit post'
        })
    },
    async mounted() {
        // Fetch post
        if (!this.postId) {
            return
        }

        let post = await API.Posts.findOne({id: this.postId})

        if (!post) {
            this.$root.$emit('show_error', 'Unable to find post.')
            return
        }

        this.post = post

        console.log(post)
    }
}
</script>

<style>

</style>
