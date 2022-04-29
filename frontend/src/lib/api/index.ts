import * as api from './openapi';

api.defaults.baseUrl = `${import.meta.env.VITE_API_ENDPOINT}${api.servers.version0ApiPath}`;
api.defaults.headers = { authorization: '' };

// hasAuthorized returns true if setAuthorization has been previously called
// with a non-empty string.
export function hasAuthorized() {
	return !!api.defaults.headers.authorization;
}

// setAuthorization sets the API's authorization header.
export function setAuthorization(auth: string) {
	api.defaults.headers.authorization = "Bearer " + auth;
}

export function errorMsg(err) {
	if (err.data && err.data.message) {
		return err.data.message;
	}
	return `${err}`;
}

export * from "./openapi";
