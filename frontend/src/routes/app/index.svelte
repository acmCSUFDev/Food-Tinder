<script>
	import * as svelte from 'svelte';
	import * as api from '$lib/api';

	let username = '';

	svelte.onMount(async () => {
		let self = await api.getSelf();
		username = self.display_name ?? self.username;
	});
</script>

<main>
	<p>
		{#if username}
			Logged in as {username}
		{:else}
			Loading...
		{/if}
	</p>
</main>
