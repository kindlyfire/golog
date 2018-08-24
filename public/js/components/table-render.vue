<template>
    <table class="table">
        <tr>
            <th v-for="h in _config.headings" :key="h">
                {{ h }}
            </th>
        </tr>
        <tr v-for="(row, i) in _config.data" :key="i">
            <td v-for="(col, j) in _config.order" :key="j">
                {{ row[col] }}
            </td>
        </tr>
    </table>
</template>


<script>
export default {
    props: {
        config: {
            headings: Array,
            data: Array,
            order: Array
        }
    },
    data: function() {
        return {}
    },
    computed: {
        _config: function() {
            let c = this.config
            if ((!c.order || c.order == []) && c.data.length > 0) {
                c.order = Object.keys(c.data[0])
            }
            console.log(c)
            return c
        }
    }
}
</script>
