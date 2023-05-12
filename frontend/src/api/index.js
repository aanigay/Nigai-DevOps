import axios from "axios";

axios.defaults.withCredentials = true

const loginConfig = {
    baseURL: 'http://localhost:8001',
    headers: {
        'Content-Type': 'application/json',
    }
}

export const LoginAPIInstance = axios.create(loginConfig)

const postsConfig = {
    baseURL: 'http://localhost:8003',
    headers: {
        'Content-Type': 'application/json',
    }
}

export const PostsAPIInstance = axios.create(postsConfig)

const groupConfig = {
    baseURL: 'http://localhost:8002',
    headers: {
        'Content-Type': 'application/json',
    }
}

export const GroupAPIInstance = axios.create(groupConfig)

const commsConfig = {
    baseURL: 'http://localhost:8004',
    headers: {
        'Content-Type': 'application/json',
    }
}

export const CommsAPIInstance = axios.create(commsConfig)