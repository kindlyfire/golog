<template>
    <div class="panel-container">
        <div class="posts">
            <table class="table">
                <tr>
                    <th>Title</th>
                    <th>Status</th>
                    <th>CommentStatus</th>
                    <th>CommentCount</th>
                    <th>Actions</th>
                </tr>
                <tr v-for="(row) in posts" :key="row.ID">
                    <td>{{ row.Title }}</td>
                    <td>{{ row.Status }}</td>
                    <td>{{ row.CommentStatus }}</td>
                    <td>{{ row.CommentCount }}</td>
                    <td>
                        <router-link :to="'/posts/edit/' + row.ID">Edit</router-link>
                    </td>
                </tr>
            </table>
        </div>
    </div>
</template>

<script>
import API from '@/api'
import TableRender from '@/components/table-render.vue'

export default {
    beforeCreate() {
        this.$root.$emit('titlebar_change', {
            title: "Posts",
            loading: true,
            button: {
                text: "Create",
                callback: () => this.$router.push('/posts/create')
            }
        })
    },
    data() {
        return {
            posts: []
        }
    },
    async mounted() {
        this.posts = await API.Posts.list()

        this.$root.$emit('titlebar_change', {
            loading: false
        })
    },
    components: {
        TableRender
    }
}
</script>

<style>

</style>
