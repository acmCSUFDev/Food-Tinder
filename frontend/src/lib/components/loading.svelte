<script>
	import Spinner from '$lib/reusable/spinner.svelte';
</script>

<div class="spin-logo">
	<Spinner slow />
	<img src="/static/bobaBub.png" alt="Mascot of Food Tinder, Boba Bub" />
</div>

<style>
	.spin-logo {
		position: relative;
		width: 100%;
		height: 100%;

		display: flex;
		justify-content: center;
		align-items: center;
	}

	.spin-logo * {
		position: absolute;
	}

	.spin-logo :global(.spinner) {
		--spinner-fore-color: var(--accent-foreground);
		--spinner-back-color: rgba(255, 255, 255, 0.45);

		width: 64px;
		height: 64px;
		border-width: 10px;
		box-shadow: 0 0 24px -2px rgba(0, 0, 0, 0.15);
	}

	.spin-logo :global(.spinner),
	.spin-logo img {
		width: 64px;
		height: 64px;
	}

	.spin-logo img {
		margin: 10px;
	}
</style>
