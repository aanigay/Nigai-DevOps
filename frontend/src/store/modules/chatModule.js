export default {
    namespaced: true,
    state: {
        //messages: JSON.parse(localStorage.getItem('messages')) || []
        messages: []
    },
    getters: {
        getMessages: (state) => state.messages
    },
    mutations: {
        SOCKET_msg(state, message) {
            console.log('message received', message)

            state.messages.push(JSON.parse(message))
            console.log(state.messages)
            localStorage.setItem('messages', JSON.stringify(state.messages))
        },
    },
}