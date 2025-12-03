import {service} from "@/api/index";


export interface QueryUserInfoResponse
{
    code: number;
    msg: string;
    data: UserInfo;
}

export interface UserInfo {
    uid : number;
    nickname : string;
    avatar : string;
    role : number;
    abstract : string
}


export interface QueryUserInfoRequest {}


//查询当前登录用户信息， 路径前面需要有 反斜杠，否则地址是相对于当前路径的相对地址
/**
 * 查询当前登录用户的用户信息
 * @param request
 */
export function queryUserInfo(request: QueryUserInfoRequest) : Promise<QueryUserInfoResponse> {
    return service.get('/api/user/user_info', {
        params: request // 【在这里传递参数】
    });
}


export interface UpdateUserInfoRequest
{
    uid : number;
    nickname? : string;
    avatar? : string ;
    abstract? : string
}

export interface UpdateUserInfoResponse
{
    code: number;
    msg: string;
    data: boolean;
}

//更新用户信息（昵称、头像等）
export function updateUserInfo(request : UpdateUserInfoRequest) : Promise<UpdateUserInfoResponse> {
    return service.post('/api/user/update_user_info', request)
}


export interface SearchUserRequest
{
    keyword? : string
}

export interface SearchUserResponse
{
    list: UserInfo[]

}
// 根据uid或者昵称搜索用户
export function searchUser(request : SearchUserRequest) : Promise<SearchUserResponse> {
    return service.get('/api/user/user_search', {
        params: request
    });
}

export interface ApplyFriendRequest
{
    uid : number;
    nickname? : string;
}

export interface ApplyFriendResponse
{
    code: number;
    msg: string;
    data: boolean;
}
//申请成为好友
export function applyFriend(request : ApplyFriendRequest) : Promise<ApplyFriendResponse> {
    return service.post('/api/user/apply_friend', request)
}


export interface FriendListResponse
{
    list: UserInfo[]

}
//我的好友列表
export function queryMyFriendList() : Promise<FriendListResponse> {
    return service.get('/api/user/friend_list')
}


export interface FriendApplyListResponse {
    code: number
    msg: string
    list: FriendApplyItem[]
}

export interface FriendApplyItem {
    friend_verify_model_id: number
    uid: number
    nickname: string
    avatar: string
    status : number
}
//查看我的好友申请验证列表
export function getFriendApplyList() : Promise<FriendApplyListResponse> {
    return service.get('/api/user/friend_apply_list')
}


export interface ResponseFriendApplyRequest {
    id : number
    status : number
}

export interface ResponseFriendApplyResponse {
    code: number
    msg: string
    data: boolean
}
//同意或者拒绝好友申请
export function handleResponseFriendApply(request : ResponseFriendApplyRequest) : Promise<ResponseFriendApplyResponse> {
    return service.put('/api/user/response_friend_apply', request)
}



