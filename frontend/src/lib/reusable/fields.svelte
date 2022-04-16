<fieldset>
	<slot />
</fieldset>

<style>
	fieldset {
		--gap: 24px;

		margin: 0;
		padding: 0;
		border: none;

		display: flex;
		flex-direction: column;

		margin-top: var(--gap);
		gap: var(--gap);
	}
</style>
