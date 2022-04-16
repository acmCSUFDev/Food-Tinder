<script>
	import InputFieldComponent from '../lib/reusable/input-field.svelte';
	import Buttons from '../lib/reusable/button.svelte';

	import Logo from '../lib/components/logo.svelte';
	// import * as api from '../lib/api';

	let loginErrorMsg = '';

	let logInButton = {
		btnType: 'login',
		btnContent: 'Log In',
		btnRoute: '/requirements',
		loginInfo: '',
		passInfo: ''
	};
	let signUpButton = {
		btnType: 'signup',
		btnContent: 'Sign Up',
		btnRoute: '/signup'
	};

	function loginError(e) {
		loginErrorMsg = e.detail;
	}

	// function onMount(async () => {
	//  	const a = await api.login('joe', 'mama');
	//  	console.log(a);
	//  });
</script>

<main class="background">
	<div class="loginComponents">
		<Logo />
		<InputFieldComponent placeholderText={'Username'} bind:value={logInButton.loginInfo} />
		<InputFieldComponent placeholderText={'Password'} bind:value={logInButton.passInfo} />
		<Buttons {...logInButton} on:failedLogin={loginError} />

		<Buttons {...signUpButton} />
		<!-- include error messages upon failed login -->
		{#if loginErrorMsg !== ''}
			<p class="error">Error: {loginErrorMsg}</p>
		{/if}
	</div>
</main>

<style>
	main.background {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}

	.loginComponents {
		display: flex;
		flex-direction: column;
		gap: 28px;
	}
</style>
