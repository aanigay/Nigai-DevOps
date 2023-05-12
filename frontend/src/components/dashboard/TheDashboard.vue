<script>
import {GroupApi} from "@/api/groupApi";
import {AuthApi} from "@/api/authApi";

export default {
    data() {
        return {
            users: [],
        }
    },
    created() {
        this.getAllUsers()
    },
    methods: {
        getAllUsers() {
            AuthApi.getAll().then((res) => {
                this.users = res.data;
            })
        },
        changeRole(id, role) {
            role = role === 'editor' ? 'moderator' : 'editor'
            GroupApi.changeRoleById(id, role).then((res) => {
                console.log('role changed', res)
            })
        }
    }

}
</script>

<template>
    <v-container>
        <v-card v-for="user in users" :key="user.id" class="ma-5" width="600">
            <v-row>
                <v-col>
                    <v-card-title>
                        {{ user.name }}
                    </v-card-title>
                    <v-card-subtitle>
                        {{ user.role }}
                    </v-card-subtitle>
                </v-col>
                <v-col>
                    <v-switch v-if="user.role !== 'admin' && this.$store.getters['authModule/getRole'] === 'admin'"
                              true-value="moderator"
                              false-value="editor"
                              v-model="user.role"
                              hide-details
                              inset
                              :label="user.role"
                              @click="changeRole(user.id, user.role)"

                    ></v-switch>
                </v-col>
            </v-row>
        </v-card>
    </v-container>
</template>

<style scoped>

</style>