<script>
	import InputField from '$lib/reusable/input-field.svelte';
	import Loading from '$lib/components/loading.svelte';
	import Button from '$lib/reusable/button.svelte';
	import Fields from '$lib/reusable/fields.svelte';
	import Error from '$lib/reusable/error.svelte';
	import Card from '$lib/reusable/card.svelte';
	import Logo from '$lib/reusable/logo.svelte';
	import * as api from '$lib/api';
	import * as globals from '$lib/globals';
	import * as navigation from '$app/navigation';
	import { onMount } from 'svelte';

	let loginErrorMsg = '';
	let username = '';
	let password = '';
	let loading = true;

	function login() {
		loading = true;

		api
			.login(username, password)
			.then((session) => {
				loading = false;
				globals.token.set('Bearer ' + session.token);
				navigation.goto('/app');
			})
			.catch((err) => {
				loading = false;
				console.error(err);
				loginErrorMsg = api.errorMsg(err);
			});
	}

	onMount(() => {
		loading = false;
	});
</script>

<main class="background centered">
	{#if loading}
		<Loading />
	{:else}
		<Card seamless>
			<Logo>
				<img class="logo" src="/static/bobaBub.png" alt="Mascot of Food Tinder, Boba Bub" />
				<h2>Food Tinder</h2>
			</Logo>

			<Fields>
				<Error box msg={loginErrorMsg} />

				<InputField placeholderText="Username" bind:value={username} />
				<InputField placeholderText="Password" bind:value={password} password />
				<Button suggested onclick={login} disabled={username == '' || password == ''}>
					Log In
				</Button>
				<Button secondary href="/signup">Register</Button>
			</Fields>
		</Card>
	{/if}
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
