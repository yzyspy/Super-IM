export interface UserInfo {
    user_id: string;
    user_name: string;
    role: number;
    token: string;
    avatar: string;
    abstract: string;
}


export function parseToken(token: string): UserInfo {
    let tokenParts = token.split('.');
    if (tokenParts.length !== 3) {
        throw new Error('Invalid token format');
    }
    let header = JSON.parse(atob(tokenParts[0]));
    let payload = JSON.parse(atob(tokenParts[1]));
    let signature = tokenParts[2];

    return {
        user_id: payload.user_id,
        user_name: payload.username,
        role: 0,
        token: token,
        avatar: '',
        abstract: ''
    }
}