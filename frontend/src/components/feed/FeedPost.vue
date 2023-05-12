<template>
    <v-container class=" mt-10">
    <div>
        {{date()}}
    </div>
    <v-card
            class="mx-auto mb-2"
            variant="outlined"
    >
        <v-card-item>
            <div>
                <div class="text-h6 mb-1">
                    {{ post.title }}
                </div>
                <div class="text-caption">{{ post.body }}</div>
            </div>
        </v-card-item>

        <v-card-actions>
            <v-list-item>
                <v-list-item-title>{{ username }}</v-list-item-title>
                <v-list-item-subtitle> {{ role }}</v-list-item-subtitle>
            </v-list-item>
            <v-row class="justify-end" v-if="isAbleToEdit">
                <v-dialog
                        v-model="editPostDialog"
                        width="800"
                >
                    <template v-slot:activator="{ props }">
                        <v-btn
                                v-bind="props"
                        >
                            <v-icon icon="mdi-square-edit-outline"></v-icon>
                        </v-btn>
                    </template>
                    <v-card>
                        <v-card-title>
                            <v-text-field v-model="editedTitle"></v-text-field>
                        </v-card-title>
                        <v-card-text>
                            <v-textarea
                                    v-model="editedBody"
                            ></v-textarea>
                        </v-card-text>
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn
                                    variant="text"
                                    @click="editPostDialog = false"
                            >
                                Cancel
                            </v-btn>
                            <v-btn
                                    variant="text"
                                    @click="editPost"
                            >
                                Apply
                            </v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
                <v-dialog
                        v-model="deletePostDialog"
                        width="200"
                >
                    <template v-slot:activator="{ props }">
                        <v-btn
                                v-bind="props"
                        >
                            <v-icon icon="mdi-delete-outline"></v-icon>
                        </v-btn>
                    </template>
                    <v-card class="align-center">
                        <v-card-text>
                            Are you sure?
                        </v-card-text>
                        <v-card-actions>
                            <v-btn color="primary" variant="text" @click="deletePostDialog = false">Cancel</v-btn>
                            <v-btn color="primary" variant="text" @click="deletePost">Yes</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
            </v-row>
        </v-card-actions>
    </v-card>
<!--    <v-btn variant="outlined">-->
<!--        Comments-->
<!--    </v-btn>-->
    <CommentSection :post_id="this.post.id"/>
    </v-container>
</template>
<script>
import {AuthApi} from "@/api/authApi";
import {PostsApi} from "@/api/postsApi";
import CommentSection from "@/components/comments/CommentSection.vue";

export default {
    data() {
        return {
            username: 'guest',
            role: 'guest',
            editPostDialog: false,
            deletePostDialog: false,
            editedTitle: this.post.title,
            editedBody: this.post.body,
            loggedInUsername: this.$store.getters['authModule/getName'],
            loggedInRole: this.$store.getters['authModule/getRole']
        }
    },
    created() {
        AuthApi.getUserById(this.post.user_id).then((res) => {
            console.log(res.data)
            this.username = res.data.name
            this.role = res.data.role
        })
    },
    updated() {
        AuthApi.getUserById(this.post.user_id).then((res) => {
            console.log(res.data)
            this.username = res.data.name
            this.role = res.data.role
        })
    },
    props: {
        post: {
            id: Number,
            title: String,
            body: String,
            user_id: Number,
            create_at: String,
            comments_ids: []
        },
        justCreated: Boolean
    },
    methods: {
        date() {
            let date = new Date(Date.parse(this.post.created_at))
            let mo = new Intl.DateTimeFormat('en', { month: '2-digit' }).format(date);
            let da = new Intl.DateTimeFormat('en', { day: '2-digit' }).format(date);

            let ho = new Intl.DateTimeFormat('ru', { hour: '2-digit' }).format(date);
            let mi = new Intl.DateTimeFormat('ru', { minute: '2-digit' }).format(date);

            return `${da}-${mo} ${ho}:${mi}`
            //return date
        },
        editPost() {
            this.editPostDialog = false
            PostsApi.updatePost(this.post.id, this.editedTitle, this.editedBody).then((res) => {
                console.log('try to update post', res)
                this.$emit('postEdited', {
                    title: this.editedTitle,
                    body: this.editedBody
                })
            })
        },
        deletePost() {
            this.deletePostDialog = false
            PostsApi.deletePost(this.post.id).then((res) => {
                console.log('deleted post', res)
                this.$emit('postDeleted')
            })
        }
    },
    computed: {
        isAbleToEdit() {
            return (this.$store.getters['authModule/getRole'] !== 'guest')
                && this.username === this.$store.getters['authModule/getName']
                || this.$store.getters['authModule/getRole'] === 'admin'
                || this.$store.getters['authModule/getRole'] === 'moderator'
        }
    },
    components: {
        CommentSection
    },
    emits: ['postEdited', 'postDeleted']
}
</script>
<style scoped>

</style>