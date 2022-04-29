<script>
	// import Buttons from '../resuable/button.svelte';
</script>

<main class="page">
	<div class="wrapper">
		<div class="btnContainer">
			<h3>left</h3>
			<h3>right</h3>
		</div>
		<h3>more info</h3>
	</div>
</main>

<style>
	/* h3 {
		position: absolute;
	} */

	.wrapper h3 {
		display: flex;
		justify-content: center;
	}

	.btnContainer {
		display: flex;
		justify-content: space-between;
		align-items: center;
		height: 80vh;
	}
</style>
