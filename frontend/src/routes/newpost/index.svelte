<script>
	import UploadImg from '$lib/reusable/upload-img.svelte';
	import Logo from '$lib/reusable/logo.svelte';
	let location = '';
	let description = '';

	let img1 = '';
	let img2 = '';
</script>

<main class="background centered">
	<form action="/login">
		<!-- Link This -->
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

		<div class="userInput">
			<label for="location">Location</label>
			<input type="text" name="location" bind:value={location} />
		</div>

		<div class="userInput">
			<label for="description">Description</label>
			<input type="text" name="description" bind:value={description} />
		</div>
		<div class="userInput">
			<label for="upload" class="imgUpload">Choose Image</label>
			<input type="file" name="upload" bind:value={img1} />
			{#if img1 !== ''}
				<input type="file" bind:value={img2} />
			{/if}
		</div>

		<input type="submit" />
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

	.userInput {
		display: flex;
		flex-direction: column;
		gap: 12px;
	}

	input[type='text'] {
		border-radius: var(--border-radius);
		border: black 1px solid;
		padding: 10px;
	}

	input[type='file'] {
		visibility: hidden;
	}
	/* fix above */

	.imgUpload {
		background-color: rebeccapurple;
	}
</style>
