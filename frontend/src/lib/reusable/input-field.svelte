<script>
	export let placeholderText = '';
	export let value = '';
	export let error = false;
	export let password = false;
	export let submit = false;
</script>

{#if submit}
	<input type="submit" value="Upload" />
{:else if password}
	<input placeholder={placeholderText} class:error bind:value type="password" />
{:else}
	<input placeholder={placeholderText} class:error bind:value type="text" />
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
</style>
