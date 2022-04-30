/**
 * OpenAPI schema for Food-Tinder
 * 0.0.1
 * DO NOT MODIFY - This file has been generated using oazapfts.
 * See https://www.npmjs.com/package/oazapfts
 */
import * as Oazapfts from "oazapfts/lib/runtime";
import * as QS from "oazapfts/lib/runtime/query";
export const defaults: Oazapfts.RequestOpts = {
    baseUrl: "/api/v0",
};
const oazapfts = Oazapfts.runtime(defaults);
export const servers = {
    version0ApiPath: "/api/v0"
};
export type LoginMetadata = {
    user_agent?: string;
};
export type Session = {
    username: string;
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
export type FoodCategories = {
    [key: string]: string[];
};
export type User = {
    username: string;
    display_name?: string;
    avatar: string;
    bio?: string;
};
export type FoodPreferences = {
    likes: string[];
    prefers: FoodCategories;
};
export type Self = User & FoodPreferences & {
    birthday?: string;
};
export type Id = string;
export type Post = {
    id: Id;
    username: string;
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
 * Get the list of all valid food categories and names.
 */
export function listFoods(opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: FoodCategories;
    } | {
        status: 500;
        data: Error;
    }>("/food/list", {
        ...opts
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
    }>("/users/@self", {
        ...opts
    }));
}
/**
 * Get a user by their username
 */
export function getUser(username: string, opts?: Oazapfts.RequestOpts) {
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
    }>(`/users/${username}`, {
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
        data: (Post & {
            liked: boolean;
        })[];
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
 * Create a new post
 */
export function createPost(post?: Post, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: Post;
    } | {
        status: 400;
        data: Error;
    } | {
        status: 500;
        data: Error;
    }>("/posts", oazapfts.json({
        ...opts,
        method: "POST",
        body: post
    })));
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
    }>("/posts/like", {
        ...opts
    }));
}
/**
 * Like the post
 */
export function postPostsLike(opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchText("/posts/like", {
        ...opts,
        method: "POST"
    }));
}
/**
 * Upload an asset
 */
export function uploadAsset(body?: Blob, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: string;
    } | {
        status: 400;
        data: Error;
    } | {
        status: 413;
        data: Error;
    } | {
        status: 500;
        data: Error;
    }>("/assets", oazapfts.json({
        ...opts,
        method: "POST",
        body
    })));
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
