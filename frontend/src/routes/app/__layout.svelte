<script>
	import * as svelte from 'svelte';
	import * as api from '$lib/api';
	import * as navigation from '$app/navigation';

	svelte.onMount(() => {
		if (!api.hasAuthorized()) {
			// Go to /login if the user hasn't logged in yet.
			navigation.goto('/login');
			return;
		}
	});
</script>
