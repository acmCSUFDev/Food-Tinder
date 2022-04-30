<script>
	import InputField from '$lib/reusable/input-field.svelte';
	import Button from '$lib/reusable/button.svelte';
	import Fields from '$lib/reusable/fields.svelte';
	import Upload from '$lib/reusable/upload-img.svelte';
	import UploadImg from '$lib/reusable/upload-img.svelte';

	let location = '';
	let description = '';
</script>

<main class="background centered">
	<Fields>
		Upload Post
		<InputField placeholderText="Location" bind:value={location} />
		<InputField placeholderText="Desciption" description bind:value={description} />
		<UploadImg />
		<Button suggested href="/app">Submit</Button>
	</Fields>
</main>
