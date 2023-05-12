<script>
import {AuthApi} from "@/api/authApi";
import {CommsApi} from "@/api/commsApi";

export default {
    data () {
        return {
            username: '',
            editCommentDialog: false,
            deleteCommentDialog: false,
            editedBody: this.comment.body
        }
    },
    created () {
        AuthApi.getUserById(this.comment.user_id).then((res) => {
            this.username = res.data.name
        })
    },
    props: {
        comment: {
            id: Number,
            body: String,
            user_id: Number,
            post_id: Number,
            created_at: String
        }
    },
    computed: {
        date() {
            let date = new Date(Date.parse(this.comment.created_at))
            let mo = new Intl.DateTimeFormat('en', { month: '2-digit' }).format(date);
            let da = new Intl.DateTimeFormat('en', { day: '2-digit' }).format(date);

            let ho = new Intl.DateTimeFormat('ru', { hour: '2-digit' }).format(date);
            let mi = new Intl.DateTimeFormat('ru', { minute: '2-digit' }).format(date);

            return `${da}-${mo} ${ho}:${mi}`
            //return date
        },
        isAbleToEdit() {
            return (this.$store.getters['authModule/getRole'] !== 'guest')
                && this.comment.user_id === this.$store.getters['authModule/getId']
                || this.$store.getters['authModule/getRole'] === 'admin'
                || this.$store.getters['authModule/getRole'] === 'moderator'
        }
    },
    methods: {
        deleteComment() {
            this.deleteCommentDialog = false
            CommsApi.deleteComment(this.comment.id).then((res) => {
                console.log('comment deleted', res)
                this.$emit('commentDeleted')
            })

        },
        editComment() {
            this.editCommentDialog = false
            CommsApi.editComment(this.comment.id, this.editedBody).then((res) => {
                console.log('comment edited', res)
                this.$emit('commentEdited', this.editedBody)
            })

        }
    },
    emits: ['commentEdited', 'commentDeleted']
}
</script>

<template>
    <v-card>
        <v-card-item>
            <v-row>
                <v-col>

                    <div>
                        <div class="text-overline mb-1 border-b">
                            {{ username }}
                        </div>
                        <div class="text-caption">
                            {{ comment.body }}
                        </div>
                    </div>
                </v-col>
                <v-col cols="2">
                    <div class="float-start d-inline-flex align-content-end text-grey-darken-1 ">
                        {{date}}
                    </div>
                </v-col>

            </v-row>
        </v-card-item>

        <v-card-actions>
            <v-row class="justify-end  d-inline-flex" v-if="isAbleToEdit">
                <v-dialog
                    v-model="editCommentDialog"
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
                        <v-card-text>
                            <v-textarea
                                v-model="editedBody"
                            ></v-textarea>
                        </v-card-text>
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn
                                variant="text"
                                @click="editCommentDialog = false"
                            >
                                Cancel
                            </v-btn>
                            <v-btn
                                variant="text"
                                @click="editComment"
                            >
                                Apply
                            </v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
                <v-dialog
                    v-model="deleteCommentDialog"
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
                            <v-btn color="primary" variant="text" @click="deleteCommentDialog = false">Cancel</v-btn>
                            <v-btn color="primary" variant="text" @click="deleteComment">Yes</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
            </v-row>
        </v-card-actions>
    </v-card>
</template>

<style scoped>

</style>