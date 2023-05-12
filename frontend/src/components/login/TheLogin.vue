<template>
    <v-sheet width="300" class="mx-auto">
        <v-form fast-fail>
            <v-text-field
                v-model="name"
                label="Username"
            ></v-text-field>

            <v-text-field
                type="password"
                v-model="password"
                label="Password"
            ></v-text-field>

            <v-btn type="submit" block class="mt-2" @submit.prevent @click="onLogin">Login</v-btn>
            <v-btn type="submit" color="#80CBC4" block class="mt-2" @submit.prevent @click="onRegister">Register</v-btn>
        </v-form>
    </v-sheet>
</template>

<script>
export default {
    data: () => ({
        name: '',
        password: '',
        rules: [
            value => {
                if (value) return true

                return 'You must enter a first name.'
            },
        ],
    }),
    methods: {
        onLogin() {
            this.$store.dispatch('authModule/onLogin', {
                name: this.name,
                password: this.password
            }).then(() => {
                this.$router.push('/chat');
            })
        },
        onRegister() {
            this.$store.dispatch('authModule/onRegister', {
                name: this.name,
                password: this.password
            }).then(() => {
                this.$router.push('/chat');
            })
        }
    }
}
</script>

<style scoped>

</style>