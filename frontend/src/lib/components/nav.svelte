<script>
	import Button from '$lib/reusable/button.svelte';
</script>

<nav>
	<div class="left">
		<h1>Food Tinder</h1>
	</div>
	<div class="right">
		<!-- Create new post button -->
		<Button ternary href="/newpost"><ion-icon name="add-circle-outline" class="right" /></Button>
		<!-- Log out button -->
		<Button ternary href="/logout"><ion-icon name="log-out-outline" /></Button>
	</div>
</nav>

<style>
	nav {
		background-color: white;
		box-shadow: 0 0 10px -6px rgba(0, 0, 0, 0.55), 0 5px 20px -8px rgba(0, 0, 0, 0.35);
		position: fixed;
		width: 100%;
		height: 75px;
		display: flex;
		justify-content: space-around;
		align-items: center;
	}

	h1 {
		margin: 0;
		color: #1d1d1d;
		transition: all 300ms;
	}

	h1:hover {
		color: var(--accent-foreground);
	}

	.right {
		display: flex;
		gap: 20px;
	}
</style>
