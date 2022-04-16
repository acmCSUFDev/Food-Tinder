<main class="background">
	<div class=""
	<div class="spinner" />
	<h3 class="loggingIn">Logging in</h3>
</main>

<style>
	main {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;

		padding: 12px;
		background-color: rgba(255, 255, 255, 0.25);
	}

	.spinner {
		--spinner-fore-color: var(--accent-foreground);
		--spinner-back-color: rgba(255, 255, 255, 0.25);

		width: 64px;
		height: 64px;
		border-width: 8px;
	}

	.loggingIn {
		margin-top: 1em;
		color: var(--accent-foreground);
	}
</style>
