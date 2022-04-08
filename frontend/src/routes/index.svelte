<script>
	import InputFieldComponent from '../lib/resuable/input-field.svelte';
	import Buttons from '../lib/resuable/button.svelte';
	import BobaBub from '../lib/components/bobaBub.svelte';

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
		btnRoute: '/signup-page'
	};

	function loginError(e) {
		loginErrorMsg = e.detail;
	}
</script>
<div class="background">
	<div class="loginComponents">
		<BobaBub />
		<InputFieldComponent placeholderText={'Username'} bind:value = {logInButton.loginInfo}/>
		<InputFieldComponent placeholderText={'Password'} bind:value = {logInButton.passInfo}/>
		<Buttons {...logInButton} on:failedLogin={loginError}/>
		<Buttons {...signUpButton} />
		<!-- include error messages upon failed login -->
		{#if loginErrorMsg !== ''}
			<p class="error">Error: {loginErrorMsg}</p>
		{/if}
	</div>
</div>

<style>
	.loginComponents {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}

	.error {
		color: #ff0000;
		font-weight: 700;
	}


	.background {
		padding-top: 12vh;
		box-sizing: border-box;
		background: linear-gradient(
			to bottom,
			rgba(247, 168, 184, 1) 0%,
			rgba(234, 171, 217, 1) 25%,
			rgba(200, 181, 245, 1) 50%,
			rgba(147, 194, 255, 1) 75%,
			rgba(85, 205, 252, 1) 100%
		);
		height: 100vh;
	}

	@media (max-width: 3000px) {
		.loginComponents {
			gap: 1vw;
		}
	}
	@media (max-width: 900px) {
		.loginComponents {
			gap: 6vw;
		}
	}
</style>
