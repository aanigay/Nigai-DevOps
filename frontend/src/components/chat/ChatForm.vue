<template>
    <v-row xs12>
        <v-text-field class="" label="Введите сообщение" outline v-model="text" @keydown.enter="send(text)"/>
    </v-row>
</template>

<script>
import { useStore } from "vuex";

export default {
    setup() {
        const store = useStore()

        function getName() {
            return store.getters["authModule/getName"]
        }

        return {
            getName
        }
    },
    data: () => ({
        text: ''
    }),
    props: {
        username: String
    },
    methods: {
        send(message) {
            this.$socket.emit('msg', JSON.stringify({
                text: message,
                name: this.getName(),
                sent_at: Date.now()
            }))
            this.text = ''
        }
    }
}
</script>