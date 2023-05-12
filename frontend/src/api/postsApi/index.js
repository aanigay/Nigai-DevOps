import {PostsAPIInstance} from "@/api";

export const PostsApi = {
    getPost(id) {
        const url = '/post';
        return PostsAPIInstance.get(url,{
            params: {
                id: id
            }
        })
    },
    getAllPosts() {
        const url = '/posts';
        return PostsAPIInstance.get(url)
    },
    updatePost(id, title, body) {
        const url = '/post'
        const data = {id, title, body}
        return PostsAPIInstance.put(url, data)
    },
    createPost(user_id, title, body) {
        const url = '/post'
        const data = {user_id, title, body}
        return PostsAPIInstance.post(url, data)
    },
    deletePost(id) {
        const url = '/post'
        return PostsAPIInstance.delete(url, {
            params: {
                id: id
            }
        })
    }
}