import {service} from "@/api/index";


export interface QueryUserInfoResponse
{
    uid : number;
    nickname : string;
    avatar : string;
    role : number;

}

export interface QueryUserInfoRequest {

}


//查询当前登录用户信息
export function queryUserInfo(request: QueryUserInfoRequest) : Promise<QueryUserInfoResponse> {
    return service.post('api/user/user_info', request)
}