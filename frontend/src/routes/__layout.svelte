<script>
	import * as api from '$lib/api';
	import * as svelte from 'svelte';
	import * as globals from '$lib/globals';
	import * as navigation from '$app/navigation';

	svelte.onMount(async () => {
		try {
			if (api.hasAuthorized() && (await api.getSelf())) {
				navigation.goto('/app');
				return;
			}
		} catch (err) {
			globals.error.set(api.errorMsg(err));
		}

		navigation.goto('/login');
	});
</script>

<slot />
