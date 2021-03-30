/* Do not change, this code is generated from Golang structs */


export class API_CreateResp {
    success: boolean;
    backupKey?: string;
    errors?: any;

    static createFrom(source: any = {}) {
        return new API_CreateResp(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.success = source["success"];
        this.backupKey = source["backupKey"];
        this.errors = source["errors"];
    }
}
export class API_CreateRequest {
    with: string;
    token: string;

    static createFrom(source: any = {}) {
        return new API_CreateRequest(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.with = source["with"];
        this.token = source["token"];
    }
}
export class API_CreateA10Token {
    userName: string;
    email: string;
    password: string;
    captcha: string;

    static createFrom(source: any = {}) {
        return new API_CreateA10Token(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.userName = source["userName"];
        this.email = source["email"];
        this.password = source["password"];
        this.captcha = source["captcha"];
    }
}