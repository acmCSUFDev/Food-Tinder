<script>
	import Loading from '$lib/components/loading.svelte';
	import * as localStorage from '$lib/local-storage.js';
	import  { onMount } from 'svelte';
	import * as api from '$lib/api';
	import * as navigation from '$app/navigation';

	onMount(() => {
		let token = localStorage.get('token');
		if (token) {
			api.setAuthorization(localStorage.get('token'));
			navigation.goto('/app');
		}
	});
</script>

<main class="background centered">
	<Loading />
</main>
