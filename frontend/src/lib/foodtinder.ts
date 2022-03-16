/**
 * OpenAPI schema for Food-Tinder
 * 0.0.1
 * DO NOT MODIFY - This file has been generated using oazapfts.
 * See https://www.npmjs.com/package/oazapfts
 */
import * as Oazapfts from "oazapfts/lib/runtime";
import * as QS from "oazapfts/lib/runtime/query";
export const defaults: Oazapfts.RequestOpts = {
    baseUrl: "/",
};
const oazapfts = Oazapfts.runtime(defaults);
export const servers = {};
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
/**
 * Log in using username and password
 */
export function login(username: string, password: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: Session;
    } | {
        status: 401;
        data: Error;
    }>(`/login${QS.query(QS.form({
        username,
        password
    }))}`, {
        ...opts,
        method: "POST"
    });
}
