/* Do not change, this code is generated from Golang structs */


export class GEN_Meta {
    alt: string;
    width: number;
    height: number;

    static createFrom(source: any = {}) {
        return new GEN_Meta(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.alt = source["alt"];
        this.width = source["width"];
        this.height = source["height"];
    }
}
export class GEN_Media {
    fileType?: number;
    meta: GEN_Meta;
    url: string;

    static createFrom(source: any = {}) {
        return new GEN_Media(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.fileType = source["fileType"];
        this.meta = this.convertValues(source["meta"], GEN_Meta);
        this.url = source["url"];
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}
export class GEN_Post {
    id: number;
    userId: number;
    created?: number;
    published: number;
    updated: number;
    content: string;
    media: boolean;
    mediaList?: GEN_Media[];

    static createFrom(source: any = {}) {
        return new GEN_Post(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.id = source["id"];
        this.userId = source["userId"];
        this.created = source["created"];
        this.published = source["published"];
        this.updated = source["updated"];
        this.content = source["content"];
        this.media = source["media"];
        this.mediaList = this.convertValues(source["mediaList"], GEN_Media);
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}
export class GEN_FollowStruct {
    following: boolean;
    requested: boolean;

    static createFrom(source: any = {}) {
        return new GEN_FollowStruct(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.following = source["following"];
        this.requested = source["requested"];
    }
}
export class GEN_User {
    id: number;
    userName: string;
    name: string;
    email?: string;
    avatar: boolean;
    avatarUrl?: string;
    bio?: string;
    selfFollow?: GEN_FollowStruct;
    userFollow?: GEN_FollowStruct;
    followers: number;
    type: number;
    created?: number;

    static createFrom(source: any = {}) {
        return new GEN_User(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.id = source["id"];
        this.userName = source["userName"];
        this.name = source["name"];
        this.email = source["email"];
        this.avatar = source["avatar"];
        this.avatarUrl = source["avatarUrl"];
        this.bio = source["bio"];
        this.selfFollow = this.convertValues(source["selfFollow"], GEN_FollowStruct);
        this.userFollow = this.convertValues(source["userFollow"], GEN_FollowStruct);
        this.followers = source["followers"];
        this.type = source["type"];
        this.created = source["created"];
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}
export class GEN_AuthRequest {
    userName: string;
    authkey: string;
    remember: boolean;
    captcha: string;

    static createFrom(source: any = {}) {
        return new GEN_AuthRequest(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.userName = source["userName"];
        this.authkey = source["authkey"];
        this.remember = source["remember"];
        this.captcha = source["captcha"];
    }
}