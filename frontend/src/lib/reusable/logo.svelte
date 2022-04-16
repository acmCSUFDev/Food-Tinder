<div>
	<slot />
</div>

<style>
	div {
		--margin: 35px;

		background-color: rgba(255, 255, 255, 0.55);
		border-radius: var(--border-radius);
		padding: 15px;
		text-align: center;
		color: var(--accent-foreground);

		padding: var(--margin) 0;
		border-radius: 20px;

		display: flex;
		flex-direction: column;
		align-items: center;
		gap: var(--margin);
	}
</style>
