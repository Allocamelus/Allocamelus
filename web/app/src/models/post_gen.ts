/* Do not change, this code is generated from Golang structs */


export class Post {
    id: number;
    userId: number;
    created?: number;
    published: number;
    updated: number;
    content: string;
    media: boolean;

    static createFrom(source: any = {}) {
        return new Post(source);
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
    }
}
export class List {
    posts: Post[];

    static createFrom(source: any = {}) {
        return new List(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.posts = this.convertValues(source["posts"], Post);
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