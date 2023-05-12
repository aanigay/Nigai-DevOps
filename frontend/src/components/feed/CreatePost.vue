<script>

import { PostsApi } from "@/api/postsApi";
export default {
    data() {
        return {
            isLoggedIn: this.$store.getters['authModule/isLoggedIn'],
            user_id: this.$store.getters['authModule/getId'],
            title: '',
            body: ''
        }
    },
    methods: {
        createPost() {
            console.log('try create post', {
                user_id: this.user_id,
                title: this.title,
                body: this.body
            })
            PostsApi.createPost(this.user_id, this.title, this.body).then((res) => {
                console.log('post created', res)
                this.$emit('createdPost')
                this.title = ''
                this.body = ''
            })
        }
    },
    emits: ['createdPost']
}
</script>

<template>
    <v-container fluid>
    <v-text-field clearable v-model="this.title" label="Title"></v-text-field>
    <v-textarea clearable v-model="this.body" label="Body"></v-textarea>
    <v-btn variant="outlined" :disabled="!isLoggedIn" @click="createPost">Create Post</v-btn>
    </v-container>
</template>

<style scoped>

</style>