/* Do not change, this code is generated from Golang structs */


export class API_AuthResp {
    success: boolean;
    userId?: number;
    error?: string;
    captcha?: string;

    static createFrom(source: any = {}) {
        return new API_AuthResp(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.success = source["success"];
        this.userId = source["userId"];
        this.error = source["error"];
        this.captcha = source["captcha"];
    }
}
export class API_AuthRequest {
    with: string;
    token: string;

    static createFrom(source: any = {}) {
        return new API_AuthRequest(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.with = source["with"];
        this.token = source["token"];
    }
}
export class API_AuthA10Token {
    userName: string;
    password: string;
    remember: boolean;
    captcha: string;

    static createFrom(source: any = {}) {
        return new API_AuthA10Token(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.userName = source["userName"];
        this.password = source["password"];
        this.remember = source["remember"];
        this.captcha = source["captcha"];
    }
}