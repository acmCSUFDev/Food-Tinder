<!-- Copy-pasted from miniCSS. -->
<script>
	export let slow = false;
</script>

<div class="spinner" class:slow />

<style>
	.spinner {
		--spinner-back-color: #ddd;
		--spinner-fore-color: #555;

		display: inline-block;
		border: 0.25rem solid var(--spinner-back-color);
		border-left: 0.25rem solid var(--spinner-fore-color);
		border-radius: 50%;
		width: 1.25rem;
		height: 1.25rem;
		animation: spinner-donut-anim 1.2s linear infinite;
	}

	.spinner.slow {
		animation: spinner-donut-anim 2.4s linear infinite;
	}

	@keyframes spinner-donut-anim {
		0% {
			transform: rotate(0deg);
		}
		100% {
			transform: rotate(360deg);
		}
	}
</style>
