import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
// 引入ElementPlus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// 引入axios
import axios from "axios";
axios.defaults.baseURL = "http://server.boychai.xyz:8088"
// axios请求拦截
// 请求拦截
axios.interceptors.request.use(config => {
    let token = localStorage.getItem("jwtToken");
    if (token) {
        config.headers['Authorization'] = token;
    }
    return config
})
// 响应拦截
axios.interceptors.response.use( response => {
        return response.data;
    },
    error => {
        if (error.response.code===400){
            localStorage.setItem('tokenStatus',error.response.data.data)
        }
        return Promise.reject(error);
    })

const app= createApp(App)

// ElementPlus加载
app.use(ElementPlus)

// axios加载
app.config.globalProperties.$axios = axios;

app.use(router).mount('#app')
// app.mount('#app')
