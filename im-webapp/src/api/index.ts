import axios from 'axios'
import {ElMessage} from "element-plus";
import {useUserInfoStore} from "@/stores";

export const service = axios.create({
    baseURL: 'http://localhost:8889', // 基础地址
    timeout: 10000, // 超时时间
    headers: { 'Content-Type': 'application/json' } // 请求头
})
//
service.interceptors.request.use(config => {
    // 添加token
//    const token = useUserInfoStore().userInfo.token
   const token : string = useUserInfoStore().getUserInfo.token
   console.log("request interceptor token = " + token)
   config.headers['token'] = token
    return config
})
//
service.interceptors.response.use(response => {
    // 处理响应数据
    if (response.status !== 200) {
       console.log('请求失败')
       ElMessage.error(response.statusText)
    }
    return response.data
})