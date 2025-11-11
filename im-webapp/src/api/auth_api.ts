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

export function doLogin(request: AuthLoginRequest) : Promise<AuthLoginResponse> {
    return service.post('api/auth/login', request)
}


export interface RegisterRequest {
    user_name: string;
    password: string;
}

export interface RegisterResponse {
    code: number;
    msg: string;
}
export function doRegister(request: RegisterRequest) : Promise<RegisterResponse> {
    return service.post('api/auth/register', request)
}