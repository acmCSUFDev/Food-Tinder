import * as store from 'svelte/store';

// writable creates a new svelte/store.writable object for the local storage
// value with the given name.
export function writable(name) {
	const value = localStorage.getItem(name);
	const writable = store.writable(value);
	writable.subscribe((v) => {
		localStorage.setdefaultItem(name, v);
	});
	return writable;
}

// get is a shorthand for getItem.
export function get(name) {
	return localStorage.getItem(name);
}
