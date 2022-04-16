<script>
	import * as svelte from 'svelte';
	import * as api from '$lib/api';
	import * as navigation from '$app/navigation';

	svelte.onMount(() => {
		if (api.hasAuthorized()) {
			navigation.goto('/app');
		} else {
			navigation.goto('/login');
		}
	});
</script>

<slot />
