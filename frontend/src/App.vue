<template>
    <v-layout>
        <v-navigation-drawer expand-on-hover rail>
            <v-list>
                <v-list-item
                        prepend-icon="mdi-account"
                        :title="store.getters['authModule/getName']"
                        :subtitle="store.getters['authModule/getRole']"
                >
                </v-list-item>
            </v-list>

            <v-divider></v-divider>

            <v-list density="compact" nav>
                <v-list-item prepend-icon="mdi-newspaper" title="Feed"
                             value="feed" to="/feed"></v-list-item>
                <v-list-item prepend-icon="mdi-chat" title="Chat"
                             value="chat" to="/chat"></v-list-item>
                <v-list-item prepend-icon="mdi-view-dashboard" title="Dashboard"
                             value="dashboard" to="/dashboard"></v-list-item>
                <v-list-item v-if="!store.getters['authModule/isLoggedIn']" prepend-icon="mdi-login" title="Login"
                             value="login" to="/login"></v-list-item>
                <v-list-item v-if="store.getters['authModule/isLoggedIn']" prepend-icon="mdi-logout" title="Logout"
                             value="logout" to="/login" @click="logout"></v-list-item>
            </v-list>
        </v-navigation-drawer>
        <router-view class="h-screen a-main"></router-view>
    </v-layout>
</template>

<script setup>

import {useStore} from "vuex";

let store = useStore();

const logout = async () => {
    await store.dispatch('authModule/onLogout');
}

</script>

<style>

.a-main {
    width: 100dvh;
    left: 100px;
}
</style>