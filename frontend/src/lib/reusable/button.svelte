<script>
	export let suggested = false;
	export let secondary = false;
	export let disabled = false;
	export let ternary = false;
	export let quaternary = false;
	export let href = '';
	export let onclick = () => {};
</script>

{#if href}
	<a class:suggested class:secondary class:ternary class:disabled {href} role="button">
		<slot />
	</a>
{:else}
	<button class:suggested class:secondary class:quaternary class:disabled on:click={onclick}>
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
		background-color: black;
		transition: all 300ms;
	}

	*.ternary:hover {
		background-color: var(--accent-foreground);
	}

	*.quaternary {
		border-color: black;
		color: black;
		background-color: white;
		transition: all 300ms;
	}

	*.quaternary:hover {
		border-color: var(--accent-foreground);
		color: var(--accent-foreground);
		background-color: rgba(255, 255, 255, 0.85);
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
