<script>
	export let placeholderText = '';
	export let value = '';
	export let error = false;
	export let password = false;
	export let description = false;
</script>

{#if password}
	<input placeholder={placeholderText} class:error bind:value type="password" />
{:else}
	<input placeholder={placeholderText} class:error class:description bind:value type="text" />
{/if}

<style>
	input {
		border: white 2px solid;
		border-radius: var(--border-radius);
		outline: none;
		padding: var(--input-padding);
		font-family: Arial, Helvetica, sans-serif;
		font-size: 18px;
		transition: ease 100ms all;
	}
	input.error {
		border-color: var(--error-foreground);
		box-shadow: 0 0 10px -2px var(--error-foreground);
	}
	input:focus {
		border-color: var(--accent-foreground);
		box-shadow: 0 0 12px -2px var(--accent-foreground);
	}
	/* comback */
	.description {
		font-size: 18px;
		display: inline-block;
		min-height: inherit;
		max-height: 600px;
	}
</style>
