<script>
	import * as globals from '$lib/globals';
	import * as navigation from '$app/navigation';

	// TODO: invalidate token on server side.
	globals.token.set('');

	navigation.goto('/login');
</script>
