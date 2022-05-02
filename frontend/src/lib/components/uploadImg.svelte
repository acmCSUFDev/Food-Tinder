<script>
	export let value = '';
	export let fileName = '';

	function removeImg() {
		value = '';
	}
</script>

{#if value !== ''}
	<!-- CHANGE SRC to /{fileName}-->
	<div class="imgUpload">
		<img src="/bobaBub.png" alt={fileName} height="100" width="100" />

		<p class="invisible" on:click={removeImg}>
			<ion-icon name="add-circle-outline" class="plus" />
		</p>
	</div>
{:else}
	<label for="upload" class="imgUpload">
		<ion-icon name="add-circle-outline" class="plus" />
		<input type="file" id="upload" bind:value />
	</label>
{/if}

<style>
	input[type='file'] {
		opacity: 0;
	}

	.imgUpload {
		border: black 2px solid;
		border-radius: var(--border-radius);
		padding: 10px;
		display: flex;
		justify-content: center;
		align-items: center;
		transition: all 300ms;
		font-size: 32px;
		height: 100px;
		width: 100px;
	}

	.imgUpload:hover {
		border-color: var(--accent-foreground);
		transform: scaleY(1.005);
	}

	.plus {
		position: absolute;
		transition: all 300ms;
	}

	.imgUpload:hover .plus {
		color: var(--accent-foreground);
		transform: rotate(180deg) scale(1.4);
	}

	.invisible {
		visibility: hidden;
		opacity: 0;
		position: absolute;
		top: 5px;
		transform: rotate(45deg);
		transition: all 300ms;
	}

	.imgUpload:hover .invisible {
		visibility: visible;
		opacity: 1;
	}

	img {
		opacity: 1;
		transition: all 300ms;
	}

	.imgUpload:hover img {
		opacity: 0.5;
	}
</style>
