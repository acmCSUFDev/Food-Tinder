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
    metadata?: LoginMetadata;
};
export type Error = {
    message: string;
};
export type User = {
    id: Id;
    name: string;
    avatar: string;
    bio?: string;
};
export type Self = User & {
    birthday: string;
};
/**
 * Log in using username and password
 */
export function login(username: string, password: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: Session;
    } | {
        status: 400;
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
export function register(username: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: Session;
    } | {
        status: 400;
        data: Error;
    }>(`/register${QS.query(QS.form({
        username
    }))}`, {
        ...opts,
        method: "POST"
    }));
}
/**
 * Get the current user
 */
export function getUsersSelf(opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: Self;
    }>("/users/self", {
        ...opts
    }));
}
/**
 * Get a user by their ID
 */
export function getUsersById(id: Id, password: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: User;
    }>(`/users/${id}${QS.query(QS.form({
        password
    }))}`, {
        ...opts
    }));
}
