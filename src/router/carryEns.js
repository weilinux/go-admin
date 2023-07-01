import axios from 'axios'
import store from "@/store";
import router from '@/router';
import {ElMessage} from "element-plus";

// const BASE_URL = 'http://www.wllinux.com:9550/'
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';

export default function request(options) {
    return new Promise((resolve, reject) => {
        const instance = axios.create({
            method: 'get',
            // baseURL: BASE_URL,
            // headers: {
            //     'Content-Type': 'application/json;charset=UTF-8',
            // },
            // timeout: 5000,
            retry: 2,
            retryDelay:100,
            // withCredentials: true,
            // responseType: 'json'
        });

        // request 请求拦截器
        instance.interceptors.request.use((config) => {
            let tokenFull = JSON.parse(localStorage.getItem('Token'));
            if ( tokenFull != null) {
                config.headers.Authorization = 'Bearer ' + tokenFull.Token
                console.log(tokenFull)
            }
            return config
        }, (error) => {
            ElMessage.error(error)
            return Promise.reject(error)
        });

        // 添加响应拦截器
        instance.interceptors.response.use(function (response) {
            let data = response.data;
            console.log(data);
            switch (data.Code) {
                case 400:
                    console.log('请求错误!');
                    break;
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
        // 请求处理
        instance(options)
            .then((res) => {
                if (res.data.code === 200 || res.data.code === 0) {
                    resolve(res);
                } else {
                    // 未登录
                    if (res.data.code === -2) {
                        router.push("/login");
                    }
                    ElMessage({ message: res.data.msg || "操作失败", type: "error", showClose: true });
                    reject(res);
                }
            })
            .catch((error) => {
                reject(error);
            });
    });
}