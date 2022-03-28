<script>
	import InputFieldComponent from '../resuable/input-field.svelte';
	import Buttons from '../resuable/button.svelte';
	import BobaBub from '../components/bobaBub.svelte';
</script>

<div>
	<!-- make white background -->
	<BobaBub />
	<InputFieldComponent placeholderText={'Username'} />
	<InputFieldComponent placeholderText={'Password'} />
	<Buttons btnType={'login'} btnContent={'Log In'} />
	<Buttons btnType={'signup'} btnContent={'Sign Up'} />
</div>

<style>
	div {
		/* padding: 12vh; */
		display: flex;
		flex-direction: column;

		justify-content: center;
		align-items: center;
	}
	@media (max-width: 3000px) {
		div {
			gap: 1vw;
		}
	}
	@media (max-width: 900px) {
		div {
			gap: 6vw;
		}
	}
</style>
