<script>
	export let box = false;
	export let msg = '';
</script>

<p class="error" class:box>
	{#if msg}
		<span class="prefix">Error:</span> {msg}
	{:else}
		<slot />
	{/if}
</p>

<style>
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
