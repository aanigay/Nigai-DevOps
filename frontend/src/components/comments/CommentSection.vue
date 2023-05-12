<script>
import TheComment from "@/components/comments/TheComment.vue";
import {CommsApi} from "@/api/commsApi";

export default {
    mounted() {
        this.getComments()
    },
    data() {
        return {
            comments: [],
            newCommentBody: ''
        }
    },
    props: {
        post_id: Number
    },
    components: {
        TheComment
    },
    methods: {
        getComments() {
            this.comments = []
            CommsApi.getByPostId(this.post_id).then((res) => {
                for (const i in res.data) {
                    CommsApi.getById(res.data[i]).then((res) => {
                        this.comments.push(res.data)
                    })
                }
            })
        },
        deleteComment() {
            this.getComments()
        },
        createComment() {
            CommsApi.createComment(this.newCommentBody, this.$store.getters['authModule/getId'], this.post_id).then((res) => {
                console.log('create comment', res)
                this.getComments()
                this.newCommentBody = ''
            })
        }
    }
}
</script>

<template>
    <v-expansion-panels>
        <v-expansion-panel
                title="Comments"
        >
            <v-expansion-panel-text>
                <div class="text-caption" v-if="!comments.length">No comments yet!</div>
                <TheComment @commentDeleted="deleteComment" @commentEdited="(body) => c.body = body" class="ma-2"
                            v-for="c in comments" :key="c.id" :comment="c"></TheComment>
                <v-text-field v-model="newCommentBody" class="mt-10" v-if="this.$store.getters['authModule/isLoggedIn']"
                              label="Comment" @keydown.enter="createComment"></v-text-field>
            </v-expansion-panel-text>
        </v-expansion-panel>
    </v-expansion-panels>
</template>

<style scoped>

</style>