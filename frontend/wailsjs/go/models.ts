export namespace commands {
	
	export class CreateUserResult {
	    id: string;
	    name: string;
	    email: string;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateUserResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.email = source["email"];
	        this.created_at = source["created_at"];
	    }
	}

}

export namespace queries {
	
	export class GetUserResult {
	    id: string;
	    name: string;
	    email: string;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new GetUserResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.email = source["email"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}

}

