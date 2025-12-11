/**
 * 初始化 WebSocket 连接
 */

const ServerUrl = 'ws://localhost:9004/api/chat/ws/chat?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOCIsInVzZXJuYW1lIjoieXp5IiwiZXhwIjoxNzY1NTMxODQzLCJpYXQiOjE3NjU0NDU0NDMsImlzcyI6InlvdXItYXBwLW5hbWUifQ.f4Z6q7-q0lSzZW1GOk8hC3zlqkxo9VXd1ViN5pahYvs'; // 对应 Go 服务器的地址和路径
let ws : WebSocket | null = null;
export function initWebSocket() : WebSocket {
    if (ws && ws.readyState === WebSocket.OPEN) return ws;

    ws = new WebSocket(ServerUrl);

    // 连接建立成功
    ws.onopen = () => {

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
}

