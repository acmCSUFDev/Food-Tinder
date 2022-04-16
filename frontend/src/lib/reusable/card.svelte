<script>
	export let small = false;
	export let large = false;
	// seamless makes the card not draw borders and padding. This is
	// particularly useful for borrowing its max-width.
	export let seamless = false;
</script>

<div class="card" class:small class:large class:seamless>
	<slot />
</div>

<style>
	.card {
		width: 100%;
		margin: 16px;
		max-width: min(360px, 90%);
	}
	.card:not(.seamless) {
		/* TODO */
		border: 1px solid var(--accent-foreground);
		border-radius: 8px;

		margin: 0;
		max-width: 360px;
	}
	.card.small {
		max-width: 240px;
	}
	.card.large {
		max-width: 480px;
	}
</style>
