
import axios from 'axios'

const get = async (url) => {
    let d = await axios.get(window.API_URL + url)

    if (d.status != 200 || d.data.Success != true) {
        console.error('Request to "' + url + '" failed.', d)
        return false
    }

    return d.data.Data
}

export default {
    Posts: {
        async list() {
            let d = await get('/posts/list')

            if (d !== false) {
                return d.Posts
            }
            
            return []
        }
    }
}