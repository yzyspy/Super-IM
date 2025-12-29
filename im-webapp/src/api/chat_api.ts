import {service} from "@/api/index";

export interface QueryChatHistoryRequest {
    page: number
    limit: number
}

export interface QueryChatHistoryResponse {

}
export function queryChatHistory(req: QueryChatHistoryRequest) : Promise<QueryChatHistoryResponse> {
    return service.get('/api/chat/history', {
        params: req
    })
}

export interface UserBaseType {
    "id": number
    "nickName": string
    "avatar": string
}

export interface TextMsg {
    "content": string
}