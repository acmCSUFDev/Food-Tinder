<script>
	import * as svelte from 'svelte';

	export let box = false;
	export let msg = '';
	export let mounted = false;

	svelte.onMount(() => {
		// Prevent the box from showing up while the JS is loading by hiding it
		// by default until everything is initialized.
		mounted = true;
	});
</script>

<p class="error" class:box class:mounted>
	{#if msg}
		<span class="prefix">Error:</span> {msg}
	{:else}
		<slot />
	{/if}
</p>

<style>
	p.error:not(.mounted),
	p.error:empty {
		display: none; /* TODO: animate */
	}
	p.error {
		color: var(--error-foreground);
	}
	p.error span.prefix {
		font-weight: 700;
	}
	p.error.box {
		background-color: var(--error-background);
		color: white;

		margin: 0;
		padding: var(--input-padding);

		border: 2px solid var(--error-foreground);
		border-radius: var(--border-radius);
	}
</style>
