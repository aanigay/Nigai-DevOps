import {PostsApi} from "@/api/postsApi";

export default {
    namespaced: true,
    state: {
        posts: [{
            id: 1,
            title: "test title",
            body: "test body",
            user_id: 1,
            comments_ids: [],
            created_at: ''
        }]
    },
    getters: {
        getPosts: (state) => {
            console.log('get: posts', state.posts)
            return [{
                id: 1,
                title: "test title",
                body: "test body",
                user_id: 1,
                comments_ids: [],
                created_at: ''
            }]
        }
    },
    mutations: {
        setPosts(state, data) {
            state.posts = data
            console.log('set: posts', data)
            //localStorage.setItem('posts', JSON.stringify(data))
        }
    },
    actions: {
        onLoadPosts({commit}) {
            PostsApi.getAllPosts().then((res) => {
                commit('setPosts', res.data)
                console.log('res', res.data)
            })
        }
    }
}