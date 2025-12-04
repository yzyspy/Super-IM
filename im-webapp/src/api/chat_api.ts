import {service} from "@/api/index";

export interface QueryChatHistoryRequest {

}

export interface QueryChatHistoryResponse {

}
export function queryChatHistory(req: QueryChatHistoryRequest) : Promise<QueryChatHistoryResponse> {
    return service.get('/api/chat/history', req)
}