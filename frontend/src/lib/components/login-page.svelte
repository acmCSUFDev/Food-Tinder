<script>
	import InputFieldComponent from '../resuable/input-field.svelte';
	import Buttons from '../resuable/button.svelte';
	import BobaBub from '../components/bobaBub.svelte';
	import { fade } from 'svelte/transition';
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();

	function clickedDispatch() {
		dispatch('loginClick', true);
	}
</script>

<div class="background" transition:fade={{ duration: 2000 }}>
	<div class="loginComponents">
		<BobaBub />
		<InputFieldComponent placeholderText={'Username'} />
		<InputFieldComponent placeholderText={'Password'} />
		<div on:click={clickedDispatch}>
			<Buttons btnType={'login'} btnContent={'Log In'} />
		</div>
		<!-- make signup button work lul -->
		<Buttons btnType={'signup'} btnContent={'Sign Up'} />
	</div>
</div>

<style>
	.loginComponents {
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
