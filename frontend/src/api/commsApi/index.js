import {CommsAPIInstance} from "@/api";

export const CommsApi = {
    getById(id) {
        const url = '/comment'
        return CommsAPIInstance.get(url, {
            params: {
                id: id
            }
        })
    },

    getByPostId(postId) {
        const url = '/commentByPostId'

        return CommsAPIInstance.get(url, {
            params: {
                postId: postId
            }
        })
        //     .then((res) => {
        //     ids = res.data
        //     console.log('got these ids', ids)
        // })
        // let comments = []
        // for (const id in ids) {
        //     this.getById(id).then((res) => {
        //         comments.push(res.data)
        //         console.log('pushed comment', res.data)
        //     })
        // }
        // return comments
    },

    createComment(body, user_id, post_id) {
        const url = '/comment'
        return CommsAPIInstance.post(url, {
            body,
            user_id,
            post_id
        })
    },

    editComment(id, body) {
        const url = '/comment'
        return CommsAPIInstance.put(url, {
            id,
            body
        })
    },

    deleteComment(id) {
        const url = '/comment'
        return CommsAPIInstance.delete(url, {
            params: {
                id: id
            }
        })
    }
}