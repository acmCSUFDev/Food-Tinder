<script>
	export let suggested = false;
	export let secondary = false;
	export let ternary = false;
	export let disabled = false;
	export let href = '';
	export let onclick = () => {};
</script>

{#if href}
	<a class:suggested class:secondary class:ternary class:disabled {href} role="button">
		<slot />
	</a>
{:else}
	<button class:suggested class:secondary class:disabled on:click={onclick}>
		<slot />
	</button>
{/if}

<style>
	a {
		text-decoration: none;
	}

	* {
		border: 2px solid transparent;
		border-radius: 50px;
		padding: 10px 15px;

		display: flex;
		justify-content: center;

		font-family: Arial, Helvetica, sans-serif;
		font-size: 18px;

		background-color: #1d1d1d;
		color: #eee;

		transition: ease 100ms all;
	}

	*.secondary {
		color: var(--accent-foreground);
		background-color: white;
	}

	*.secondary:hover {
		border-color: var(--accent-foreground);
		background-color: rgba(255, 255, 255, 0.85);
	}

	*.ternary {
		background-color: #1d1d1d;
		transition: all 300ms;
	}

	*.ternary:hover {
		background-color: var(--accent-foreground);
	}

	*.suggested {
		color: white;
		background-color: var(--accent-border);
		border-color: var(--accent-foreground);
		box-shadow: 0 0 10px 0px var(--accent-foreground);
	}

	*.suggested:hover {
		background-color: var(--accent-background);
	}

	*.disabled {
		pointer-events: none;
		opacity: 0.55;
		cursor: not-allowed;
		box-shadow: none;
	}
</style>
