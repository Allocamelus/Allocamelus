/* Do not change, this code is generated from Golang structs */


export class User {
    id: number;
    userName: string;
    name: string;
    email?: string;
    avatar: boolean;
    avatarUrl?: string;
    bio?: string;
    likes: number;
    created?: number;

    static createFrom(source: any = {}) {
        return new User(source);
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
        this.likes = source["likes"];
        this.created = source["created"];
    }
}
export class Session {
    loggedIn: boolean;
    userId: number;
    userName: string;
    perms: number;
    notNew: boolean;

    static createFrom(source: any = {}) {
        return new Session(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.loggedIn = source["loggedIn"];
        this.userId = source["userId"];
        this.userName = source["userName"];
        this.perms = source["perms"];
        this.notNew = source["notNew"];
    }
}