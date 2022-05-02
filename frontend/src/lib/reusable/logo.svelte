<script>
	export let intro = false;
</script>

<div class:intro>
	<slot />
</div>

<style>
	div {
		--margin: 35px;

		background-color: rgba(255, 255, 255, 0.55);
		border-radius: var(--border-radius);

		text-align: center;
		color: var(--accent-foreground);

		border-radius: 20px;

		display: flex;
		flex-direction: column;
		align-items: center;
		gap: var(--margin);
	}

	div.intro {
		padding: 15px;
		padding: var(--margin) 0;
	}
</style>
