import {service} from "@/api/index";

export interface AuthLoginRequest {
    user_name: string;
    password: string;
}

export interface AuthLoginResponse {
    code: number;
    msg: string;
    data: {
        token: string;
    }
}


/**
 * 用户登录
 * @param request
 */

export function doLogin(request: AuthLoginRequest) : Promise<AuthLoginResponse> {
    return service.post('/api/auth/login', request)
}



export interface AuthLogoutRequest {
}

export interface AuthLogoutResponse {
    code: number;
    msg: string;
}

/**
 * 用户退出登录
 * @param request
 */
export function doLogout(request: AuthLogoutRequest) : Promise<AuthLogoutResponse> {
    return service.post('/api/auth/logout', request)
}


export interface RegisterRequest {
    user_name: string;
    password: string;
}

export interface RegisterResponse {
    code: number;
    msg: string;
}

/**
 * 用户注册
 * @param request
 */
export function doRegister(request: RegisterRequest) : Promise<RegisterResponse> {
    return service.post('/api/auth/register', request)
}