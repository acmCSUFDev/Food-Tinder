<script>
	export let placeholderText = '';
	export let value = '';
</script>

<input type="text" placeholder={placeholderText} bind:value />

<!-- <p>AHHHHHHHHHHH</p> -->
<style>
	input {
		border: white 1px solid;
		border-radius: 50px;
		padding: 10px 15px;
		font-family: Arial, Helvetica, sans-serif;
		font-size: 18px;
	}
</style>
