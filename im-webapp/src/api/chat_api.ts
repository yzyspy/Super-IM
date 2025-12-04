import {service} from "@/api/index";

export interface QueryChatHistoryRequest {
    page: number
    limit: number
}

export interface QueryChatHistoryResponse {

}
export function queryChatHistory(req: QueryChatHistoryRequest) : Promise<QueryChatHistoryResponse> {
    return service.get('/api/chat/history', req)
}