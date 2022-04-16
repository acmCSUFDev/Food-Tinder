<script>
	import InputField from '../lib/reusable/input-field.svelte';
	import Button from '../lib/reusable/button.svelte';
	import Fields from '../lib/reusable/fields.svelte';
	import Error from '../lib/reusable/error.svelte';
	import Card from '../lib/reusable/card.svelte';
	import Logo from '../lib/reusable/logo.svelte';

	let username = '';
	let password = '';
	let passwordVerify = '';
	let displayName = '';
	let errorMsg = '';

	function register() {}
</script>

<main class="background centered">
	<Card seamless>
		<Logo>
			<h2>Register</h2>
		</Logo>

		<Fields>
			<Error box msg={errorMsg} />

			<InputField placeholderText="Username" bind:value={username} />
			<InputField placeholderText="Nickname (optional)" bind:value={displayName} />
			<InputField placeholderText="Password" password bind:value={password} />
			<InputField
				placeholderText="Verify Password"
				password
				bind:value={passwordVerify}
				error={passwordVerify != password}
			/>
			<Button
				suggested
				on:click={register}
				disabled={username == '' || password == '' || passwordVerify != password}
			>
				Register
			</Button>
			<Button secondary href="/login">Back to Login</Button>
		</Fields>
	</Card>
</main>
