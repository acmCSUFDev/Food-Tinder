<script>
	import Logo from '$lib/reusable/logo.svelte';
	import InputField from '$lib/reusable/input-field.svelte';
	import Button from '$lib/reusable/button.svelte';
	import Images from '$lib/components/uploadImg.svelte';

	let btnResult = false;
	let location = '';
	let description = '';

	let img1 = '';
	let img2 = '';
	let img3 = '';
	let img4 = '';

	function nextPage() {
		if (btnResult) {
			btnResult = false;
		} else {
			btnResult = true;
		}
	}
</script>

<main class="background centered">
	<form action="/login">
		{#if !btnResult}
			<Logo>
				<img
					class="logo"
					src="/static/bobaBub.png"
					alt="Mascot of Food Tinder, Boba Bub"
					width="150"
					height="150"
				/>
				<h3>Create New Post</h3>
			</Logo>

			<InputField placeholderText="Location" bind:value={location} />
			<InputField placeholderText="Description" bind:value={description} />
			<Button suggested onclick={nextPage}>Next</Button>
		{:else}
			<h3 class="uploadTitle">Upload Images</h3>
			<div class="img-Grid">
				<Images fileName={img1} bind:value={img1} />
				<Images fileName={img2} bind:value={img2} />
				<Images fileName={img3} bind:value={img3} />
				<Images fileName={img4} bind:value={img4} />
			</div>
			<Button suggested secondary onclick={nextPage}>Go Back</Button>
			<InputField submit />
		{/if}
	</form>
</main>

<style>
	form {
		display: flex;
		flex-direction: column;
		gap: 20px;
		background-color: white;
		padding: 24px;
		border-radius: var(--border-radius);
	}

	.img-Grid {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 20px;
	}

	.uploadTitle {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 20px;
		color: var(--accent-foreground);
		display: flex;
		align-items: center;
		justify-content: center;
	}
</style>
