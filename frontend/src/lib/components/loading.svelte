<script>
	import Spinner from '../reusable/spinner.svelte';
</script>

<div class="spin-logo">
	<Spinner slow />
	<img src="/static/bobaBub.png" alt="Mascot of Food Tinder, Boba Bub" />
</div>

<style>
	.spin-logo {
		position: relative;
	}

	.spin-logo * {
		position: absolute;
		top: 0;
		left: 0;
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
