<script>
	export let btnType = '';
	export let btnContent = '';
	export let activeButton = false;
</script>

<button class="{btnType}{activeButton === true ? 'Active' : ''}">{btnContent}</button>

<style>
	button {
		border: black 1px solid;
		border-radius: 50px;
		padding: 5px 10px;
		/* min-width: 60vw; */
		width: 200px;
		display: flex;
		justify-content: center;
		font-family: Arial, Helvetica, sans-serif;
		font-size: 24px;
		font-weight: 100;
	}

	.login {
		color: white;
		background-color: rgba(168, 85, 252, 0.57);
		border: rgba(168, 85, 252, 0.57) 1px solid;
	}
	.signup {
		color: rgba(154, 72, 235, 0.57);
		background-color: rgba(255, 255, 255, 0.45);
		border: rgba(255, 255, 255, 0.45) 1px solid;
	}
	.loginActive {
		color: white;
		background-color: rgba(168, 85, 252, 0.767);
		border: rgba(168, 85, 252, 0.767) 1px solid;
	}
	.signupActive {
		color: rgba(168, 85, 252, 0.57);
		background-color: rgba(255, 255, 255, 0.596);
		border: rgba(255, 255, 255, 0.596) 1px solid;
	}
</style>
