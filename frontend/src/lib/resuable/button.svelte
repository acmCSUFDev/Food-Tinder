<script>
	import { createEventDispatcher } from 'svelte';

	export let btnType = '';
	export let btnContent = '';
	export let btnRoute = '#';
	export let loginInfo = '';
	export let passInfo = '';

	const dispatch = createEventDispatcher();

	function failedLoginDispatch() {
		if(loginInfo === '') {
			dispatch('failedLogin', 'username cannot be left blank');
		} else if (passInfo === '') {
			dispatch('failedLogin', 'password cannot be left blank');
		} else {
			dispatch('failedLogin', 'username and password cannot be left blank');
		}
	}
</script>
{#if (loginInfo !== '') && (passInfo !== '')}
	<a href="{btnRoute}">
		<button class="{btnType}">{btnContent}</button>
	</a>
{:else if (btnType === 'signup') || (btnType === 'signup2' )}
<a href="{btnRoute}">
	<button class="{btnType}">{btnContent}</button>
</a>
{:else}
<button class="{btnType}" on:click={failedLoginDispatch}>{btnContent}</button>
{/if}


<style>
	a {
		text-decoration: none;
	}
	button {
		border: black 1px solid;
		border-radius: 50px;
		padding: 10px 15px;
		/* min-width: 60vw; */
		width: 250px;
		display: flex;
		justify-content: center;
		font-family: Arial, Helvetica, sans-serif;
		font-size: 18px;
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
</style>
