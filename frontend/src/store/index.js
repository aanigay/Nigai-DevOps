
import { createStore } from "vuex";
import authModule from './modules/authModule';
import chatModule from "@/store/modules/chatModule";
import postsModule from "@/store/modules/postsModule";


export default createStore({
    modules: {
        authModule,
        chatModule,
        postsModule
    }
})