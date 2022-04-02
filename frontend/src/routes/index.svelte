<script>
	import UserConstraints from '../lib/components/requirements.svelte';
	import Login from '../lib/components/login-page.svelte';
	import SignUp from '../lib/components/signup-page.svelte';
	let loggedIn = false;
	let registered = false;

	function clickedLoginEvent(e) {
		loggedIn = e.detail;
		console.log('successfully logged in');
	}
	// function clickedSignUpEvent(e) {
	// 	registered = e.detail;
	// 	console.log('successfully signed up');
	// }
</script>

{#if loggedIn && registered}
	<UserConstraints />
{:else if !registered && !loggedIn}
	<SignUp />
{:else}
	<Login on:loginClick={clickedLoginEvent} />
{/if}
