<script>
	// import InputFieldComponent from '../resuable/input-field.svelte';
	// import Buttons from '../resuable/button.svelte';
	// import BobaBub from '../components/bobaBub.svelte';
	import { fade } from 'svelte/transition';
	// import { createEventDispatcher } from 'svelte';
	// export let activeButton = false;

	// function clickedEvent(e) {
	// 	activeButton = e.detail;
	// 	console.log(activeButton);
	// }
	// const dispatch = createEventDispatcher();

	// function clickedDispatch() {
	// 	dispatch('signupClick', true);
	// }
</script>

<div class="background" transition:fade={{ duration: 2000 }}>
	<p>
		Lorem ipsum dolor sit amet consectetur adipisicing elit. Id inventore a libero maxime aperiam
		sit error soluta voluptate earum corrupti.
	</p>
</div>

<style>
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
</style>
