<script>
	import InputField from '$lib/reusable/input-field.svelte';
	import Button from '$lib/reusable/button.svelte';
	import Fields from '$lib/reusable/fields.svelte';
	import Error from '$lib/reusable/error.svelte';
	import Card from '$lib/reusable/card.svelte';
	import Logo from '$lib/reusable/logo.svelte';

	let loginErrorMsg = '';
	let username = '';
	let password = '';

	function login() {}
</script>

<main class="background centered">
	<Card seamless>
		<Logo>
			<img class="logo" src="/static/bobaBub.png" alt="Mascot of Food Tinder, Boba Bub" />
			<h2>Food Tinder</h2>
		</Logo>

		<Fields>
			<Error box msg={loginErrorMsg} />

			<InputField placeholderText="Username" bind:value={username} />
			<InputField placeholderText="Password" bind:value={password} password />
			<Button suggested on:click={login} disabled={username == '' || password == ''}>Log In</Button>
			<Button secondary href="/signup">Register</Button>
		</Fields>
	</Card>
</main>

<style>
	img.logo {
		width: 150px;
		height: 150px;
	}

	h2 {
		color: var(--accent-foreground);
		margin: 0;
	}
</style>
