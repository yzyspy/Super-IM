export interface UserInfo {
    user_id: string;
    user_name: string;
    role: number;
}


export function parseToken(token: string): UserInfo {
    let tokenParts = token.split('.');
    if (tokenParts.length !== 3) {
        throw new Error('Invalid token format');
    }
    let header = JSON.parse(atob(tokenParts[0]));
    let payload = JSON.parse(atob(tokenParts[1]));
    let signature = tokenParts[2];

    console.log(payload);

    return {
        user_id: payload.user_id,
        user_name: payload.username,
        role: 0
    }
}