import * as store from 'svelte/store';

var known = {};

// persistent creates a new svelte/store.writable object for the local storage
// value with the given name.
export function persistent(name, def = null) {
	if (known[name]) {
		return known[name];
	}

	let value = def;
	try {
		let v = JSON.parse(localStorage.getItem(name));
		if (v) {
			value = v;
		}
	} catch (err) {
		console.error(`local storage ${name} error (${err}), resetting to default`);
	}

	const writable = store.writable(value);
	writable.subscribe((v) => {
		localStorage.setItem(name, JSON.stringify(v));
	});

	known[name] = writable;
	return writable;
}

// get is a shorthand for getItem.
export function get(name) {
	return localStorage.getItem(name);
}
