<script>
	import InputFieldComponent from '../lib/resuable/input-field.svelte';
	import Buttons from '../lib/resuable/button.svelte';
	import RegisterBox from '../lib/components/register-box.svelte'

	let signupInformation = {
		username: '',
		displayName: '',
		password: '',
		passVerify: ''
	};

	let registerButton = {
		btnType: 'signup2',
		btnContent: 'Sign Up',
		btnRoute: '/',
	};
</script>

<div class="background">
	<div class="signUpComponents">
		<div class="item-wrapper">
			<div class="input">
				<RegisterBox />
				<InputFieldComponent placeholderText={'Username'} bind:value = {signupInformation.username}/>
				<InputFieldComponent placeholderText={'Display Name (optional)'} bind:value = {signupInformation.displayName}/>
				<InputFieldComponent placeholderText={'Password'} bind:value = {signupInformation.password}/>
				<InputFieldComponent placeholderText={'Verify Password'} bind:value = {signupInformation.passVerify}/>
			</div>
			<Buttons {...registerButton}/>
		</div>
	</div>
</div>

<style>

	.item-wrapper {
		display: flex;
		flex-direction: column;
		gap: 100px; 
	}

	.input {
		display: flex;
		flex-direction: column;
		gap: 15px;
	}

	.signUpComponents {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}

	.background {
		padding-top: 12vh;
		box-sizing: border-box;
		background: linear-gradient(
			to bottom,
			rgba(85, 205, 252, 1) 0%,
			rgba(147, 194, 255, 1) 25%,
			rgba(200, 181, 245, 1) 50%,
			rgba(234, 171, 217, 1) 75%,
			rgba(247, 168, 184, 1) 100%
		);
		height: 100vh;
	}

	@media (max-width: 3000px) {
		.signUpComponents {
			gap: 1vw;
		}
	}
	@media (max-width: 900px) {
		.signUpComponents {
			gap: 6vw;
		}
	}
</style>
