<template>
    <v-container>
        <v-row>
            <v-col cols="5">
                <v-container class="mt-6 position-sticky" fluid>
                    <CreatePost @createdPost="getAllPosts" class="position-static"/>
                </v-container>
            </v-col>
            <v-spacer></v-spacer>
            <v-spacer></v-spacer>
            <v-spacer></v-spacer>
            <v-col cols="7">
                <v-container class="feed">
                    <v-virtual-scroll :items="items" height="750">
                        <template v-slot:default="{ item }">
                            <FeedPost @postDeleted="getAllPosts"
                                      @postEdited="(data) => {item.body = data.body; item.title = data.title}"
                                      :post="item"/>
                        </template>
                    </v-virtual-scroll>
                </v-container>
            </v-col>
        </v-row>
    </v-container>
</template>


<script>
import {PostsApi} from "@/api/postsApi";
import FeedPost from "@/components/feed/FeedPost.vue";
import CreatePost from "@/components/feed/CreatePost.vue"

export default {
    data() {
        return {
            items: []
        }
    },
    mounted() {
        this.getAllPosts()
    },
    components: {
        CreatePost,
        FeedPost
    },
    methods: {
        getAllPosts() {
            PostsApi.getAllPosts().then((res) => {
                console.log(res.data)
                this.items = res.data.reverse()
                this.$forceUpdate()
            })
        }

    },
    emits: ['postEdited', 'postDeleted']
}

</script>

<style scoped>
.feed {
    width: 100dvh;
    overflow-y: auto;
    height: 100%;
}
</style>