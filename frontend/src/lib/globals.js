import * as localStorage from './local-storage.js';
import * as store from 'svelte/store';
import * as api from './api';

export let token = localStorage.persistent('token');

token.subscribe((v) => {
	api.setAuthorization(v);
});

export let error = store.writable('');
