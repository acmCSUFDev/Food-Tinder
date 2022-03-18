/**
 * OpenAPI schema for Food-Tinder
 * 0.0.1
 * DO NOT MODIFY - This file has been generated using oazapfts.
 * See https://www.npmjs.com/package/oazapfts
 */
import * as Oazapfts from "oazapfts/lib/runtime";
import * as QS from "oazapfts/lib/runtime/query";
export const defaults: Oazapfts.RequestOpts = {
    baseUrl: "https://localhost/api/v0",
};
const oazapfts = Oazapfts.runtime(defaults);
export const servers = {
    version0ApiPath: ({ hostname = "localhost" }: {
        hostname: string | number | boolean;
    }) => `https://${hostname}/api/v0`
};
export type Id = string;
export type LoginMetadata = {
    user_agent?: string;
};
export type Session = {
    user_id: Id;
    token: string;
    expiry: string;
    metadata: LoginMetadata;
};
export type Error = {
    message: string;
};
export type FormError = Error & {
    form_id?: string;
};
export type User = {
    id: Id;
    name: string;
    avatar: string;
    bio?: string;
};
export type FoodPreferences = {
    likes: string[];
    prefers: {
        [key: string]: string[];
    };
};
export type Self = User & FoodPreferences & {
    birthday: string;
};
export type Post = {
    id: Id;
    user_id: Id;
    cover_hash?: string;
    images: string[];
    description: string;
    tags: string[];
    location?: string;
};
/**
 * Log in using username and password. A 401 is returned if the information is incorrect.
 */
export function login(username: string, password: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: Session;
    } | {
        status: 401;
        data: Error;
    } | {
        status: 500;
        data: Error;
    }>(`/login${QS.query(QS.form({
        username,
        password
    }))}`, {
        ...opts,
        method: "POST"
    }));
}
/**
 * Register using username and password
 */
export function register(username: string, password: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: Session;
    } | {
        status: 400;
        data: FormError;
    } | {
        status: 500;
        data: Error;
    }>(`/register${QS.query(QS.form({
        username,
        password
    }))}`, {
        ...opts,
        method: "POST"
    }));
}
/**
 * Get the current user
 */
export function getSelf(opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: Self;
    } | {
        status: 500;
        data: Error;
    }>("/users/self", {
        ...opts
    }));
}
/**
 * Get a user by their ID
 */
export function getUser(id: Id, password: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: User;
    } | {
        status: 400;
        data: FormError;
    } | {
        status: 404;
        data: FormError;
    } | {
        status: 500;
        data: Error;
    }>(`/users/${id}${QS.query(QS.form({
        password
    }))}`, {
        ...opts
    }));
}
/**
 * Get the next batch of posts
 */
export function getNextPosts({ prevId }: {
    prevId?: Id;
} = {}, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: Post[];
    } | {
        status: 400;
        data: FormError;
    } | {
        status: 500;
        data: Error;
    }>(`/posts${QS.query(QS.form({
        prev_id: prevId
    }))}`, {
        ...opts
    }));
}
/**
 * Delete the current user's posts by ID. A 401 is returned if the user tries to delete someone else's post.
 */
export function deletePost(id: Id, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
    } | {
        status: 404;
        data: FormError;
    } | {
        status: 500;
        data: Error;
    }>(`/posts/${id}`, {
        ...opts,
        method: "DELETE"
    }));
}
/**
 * Get the list of posts liked by the user
 */
export function getLikedPosts(opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: Post[];
    } | {
        status: 500;
        data: Error;
    }>("/posts/liked", {
        ...opts
    }));
}
/**
 * Get the file asset by the given ID. Note that assets are not separated by type; the user must assume the type from the context that the asset ID is from.
 */
export function getAsset(id: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: Blob;
    } | {
        status: 404;
        data: FormError;
    } | {
        status: 500;
        data: Error;
    }>(`/assets/${id}`, {
        ...opts
    }));
}
