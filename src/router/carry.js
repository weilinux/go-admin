import axios from 'axios'
import store from "@/store";
import router from "@/router";
import {ElMessage} from "element-plus";

// axios.defaults.baseURL = BASE_URL;
axios.interceptors.request.use((config) => {
    let tokenFull = JSON.parse(localStorage.getItem('Token'));
    if ( tokenFull != null) {
        config.headers.Authorization = 'Bearer ' + tokenFull.Token
        console.log(tokenFull)
    }
    return config
}, (error) => {
    ElMessage.error(error)
    return Promise.reject(error)
})

// 添加响应拦截器
axios.interceptors.response.use(function (response) {
    let data = response.data;
    console.log(data);
    switch (data.Code) {
        case 401:
            // remove token
            //store.commit()处理当前用户状态，登录状态,路由转向
            //router.replace()
            ElMessage.error("登录已过期，请重新登录!")
            store.dispatch('auth/logout');
            router.push('/login');
            // setTimeout()
            break;
    }
    return response
}, function (error) {
    ElMessage.error(error)
    return Promise.reject(error);
});

export default axios


// axios.defaults.timeout = 10000;
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8;multipart/form-data;boundary=----WebKitFormBoundaryeOOd9EoQVo6B4Tgb';
// axios.defaults.baseURL = 'http://localhost:8081/api';
// axios.defaults.baseURL = 'http://127.0.0.1:5000'
// axios.defaults.timeout = 5000  // 超时时间（毫秒）
// axios.defaults.retry = 2  // 重试次数
// axios.defaults.retryDelay = 100  // 重试之间的间隔时间（毫秒）