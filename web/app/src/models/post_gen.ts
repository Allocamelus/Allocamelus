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