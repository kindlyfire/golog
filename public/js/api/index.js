
import axios from 'axios'

const get = async (url, conf = {}) => {
    let d = await axios.get(window.API_URL + url, conf)

    if (d.status != 200 || d.data.Success != true) {
        console.error('Request to "' + url + '" failed.', d)
        return false
    }

    return d.data.Data
}

export default {
    Posts: {
        async list() {
            let d = await get('/posts')

            if (d !== false) {
                return d.Posts
            }
            
            return []
        },
        async find(params = {}) {
            let d = await get('/posts/find', {
                params
            })

            if (d !== false) {
                return d.Posts
            }

            return []
        },
        async findOne(params = {}) {
            params.limit = 1
            let posts = await this.find(params)
            return posts.length > 0 ? posts[0] : null
        }
    }
}