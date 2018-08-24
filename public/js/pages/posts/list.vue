<template>
    <div class="panel-container">
        <div class="posts">
            <table-render :config="tableData"></table-render>
        </div>
    </div>
</template>

<script>
import API from '@/api'
import TableRender from '@/components/table-render.vue'

export default {
    beforeCreate() {
        this.$root.$emit('titleBarChange', {
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
            posts: [],
            tableData: {
                headings: ['Title', 'Status', 'CommentStatus', 'CommentCount'],
                data: [],
                order: ['ID', 'Title', 'CommentStatus', 'CommentCount']
            }
        }
    },
    async mounted() {
        this.posts = await API.Posts.list()
        this.tableData.data = this.posts

        this.$root.$emit('titleBarChange', {
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
