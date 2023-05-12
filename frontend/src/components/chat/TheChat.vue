<template>
    <div class="c-wrap">
        <div class="c-chat overflow-y-auto">
            <v-virtual-scroll :items="messages" height="650">
                <template v-slot:default="{ item }">
                    <ChatMessage :name="item.name" :text="item.text" :sent_at="item.sent_at" :owner="username=== item.name"/>
                </template>
            </v-virtual-scroll>
<!--            <ChatMessage class="overflow-y-auto" v-for="m in messages" :key="m.text" :name="m.name" :text="m.text" :sent_at="m.sent_at" :owner="username===m.name"/>-->
        </div>
        <div class="c-form">
            <ChatForm  :username="username"/>
        </div>
    </div>
</template>

<script setup>
import ChatMessage from "../chat/ChatMessage.vue";
import ChatForm from "../chat/ChatForm.vue"
import {useStore} from "vuex";
import FeedPost from "@/components/feed/FeedPost.vue";
let store = useStore();
let messages = store.getters["chatModule/getMessages"] || []
let username = store.getters["authModule/getName"]

</script>

<style scoped>
.c-wrap {
    height: 100%;
    position: relative;
    overflow: hidden;
}

.c-form {
    position: absolute;
    bottom: 50px;
    left: 0;
    right: 0;
    height: 80px;
    border-radius: 10px;
}

.c-chat {
}
</style>


