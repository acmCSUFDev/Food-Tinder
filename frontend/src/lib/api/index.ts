import * as api from './openapi';

export const url = `${import.meta.env.VITE_API_ENDPOINT}${api.servers.version0ApiPath}`;

api.defaults.baseUrl = url;
api.defaults.headers = { authorization: '' };

// hasAuthorized returns true if setAuthorization has been previously called
// with a non-empty string.
export function hasAuthorized() {
	return !!api.defaults.headers.authorization;
}

// setAuthorization sets the API's authorization header.
export function setAuthorization(auth: string) {
	api.defaults.headers.authorization = auth;
}

export function errorMsg(err) {
	if (err.status == 403) {
		if (!hasAuthorized()) {
			return "missing token"
		}
		return "session expired"
	}
	if (err.data && err.data.message) {
		return err.data.message;
	}
	return `${err}`;
}

export function assetURL(id) {
	if (!id) {
		return ""
	}
	return url + "/assets/" + id;
}

export * from "./openapi";
