/**
 * 初始化 WebSocket 连接
 */

const ServerUrl = 'ws://localhost:9004/api/chat/ws/chat?token='; // 对应 Go 服务器的地址和路径
let ws : WebSocket | null = null;
export function initWebSocket(token : string) : WebSocket {
    if (ws && ws.readyState === WebSocket.OPEN) return ws;

    ws = new WebSocket(ServerUrl + token);

    // 连接建立成功
    ws.onopen = () => {
        console.log("WebSocket 已连接");
    };

    // 接收到消息
    ws.onmessage = handleWebSocketMessage;

    // 连接关闭
    ws.onclose = () => {
        console.log("WebSocket 已关闭");
    };

    // 发生错误
    ws.onerror = (err) => {
        console.error("WebSocket 发生错误:", err);
    };
    return ws;
}

/**
 * 处理 WebSocket 接收到的消息
 */
function handleWebSocketMessage(event : MessageEvent) {
    const message = event.data;
    console.log("收到消息:", message);
}

