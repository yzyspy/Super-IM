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


/**
 * 与某一个好友的历史聊天消息
 */
export interface ChatHistoryType {
    "id": number
    "sendUser": UserBaseType
    "revUser": UserBaseType
    "isMe": boolean
    "created_at": string
    "msg": {
        "type": number
        "textMsg"?: TextMsg
        "systemMsg": null
    }
}

export interface UserBaseType {
    "id": number
    "nickName": string
    "avatar": string
}

export interface TextMsg {
    "content": string
}

/**
 * websocket 接收到到某一条的用户实时发来的消息
 */
export interface MsgType {
    rev_user: UserBaseType;      // 接收者
    send_user: UserBaseType;     // 发送者
    msg: MessageContent;     // 消息主体
    create_at: string;       // ISO 8601 时间格式
}

/**
 * 消息体内容结构
 * 考虑到扩展性，将各类可选消息类型设为 null | any
 */
interface MessageContent {
    msg_type: number;
    content: string;
    image_msg: any | null;
    video_msg: any | null;
    file_msg: any | null;
    voice_msg: any | null;
    video_call_msg: any | null;
    voice_call_msg: any | null;
    witchdraw_msg: any | null; // 注意：JSON中为 witchdraw (可能为 withdraw 的拼写错误)
    reply_msg: any | null;
    quote_msg: any | null;
    at_msg: any | null;
}