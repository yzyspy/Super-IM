import axios from 'axios'
import {ElMessage} from "element-plus";
import {useUserInfoStore} from "@/stores";

export const service = axios.create({
    baseURL: '', // 基础地址, 如果使用代理，这个地方要配置位空
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
        return Promise.reject(response)
    }
    return response.data
})

export interface  BaseResponse<T> {
    code: number
    msg: string
    data: T
}

export interface ListResponse<T> {
    list: T[]
    total: number
}
