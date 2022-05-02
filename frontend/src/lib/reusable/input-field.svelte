<script>
	export let placeholderText = '';
	export let value = '';
	export let error = false;
	export let password = false;
	export let textarea = false;
	export let location = false;
	export let submit = false;
</script>

{#if submit}
	<input type="submit" value="Upload" />
{:else if textarea}
	<textarea class="desc" placeholder={placeholderText} cols="30" rows="6" bind:value />
{:else if location}
	<input placeholder={placeholderText} class:error class:location bind:value type="text" />
{:else if password}
	<input placeholder={placeholderText} class:error bind:value type="password" />
{:else}
	<input placeholder={placeholderText} class:error class:location bind:value type="text" />
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

	input.location {
		border-color: black;
	}

	input:focus,
	textarea:focus {
		border-color: var(--accent-foreground);
		box-shadow: 0 0 12px -2px var(--accent-foreground);
	}

	.desc {
		border-color: black;
		resize: none;
		/* removes ugly scrollbar for firefox */
		scrollbar-width: none;
	}
	/* removes scrollbar for every other browser */
	.desc::-webkit-scrollbar {
		display: none;
	}

	input[type='submit'] {
		background-color: var(--accent-border);
		border-color: var(--accent-foreground);
		color: white;
		box-shadow: 0 0 10px 0px var(--accent-foreground);
		max-width: 100%;
	}

	input[type='submit']:hover {
		background-color: var(--accent-background);
	}

	.location,
	textarea {
		border: white 2px solid;
		border-radius: var(--border-radius);
		outline: none;
		padding: var(--input-padding);
		font-family: Arial, Helvetica, sans-serif;
		font-size: 18px;
		max-width: 234px;
		transition: ease 100ms all;
	}
</style>
