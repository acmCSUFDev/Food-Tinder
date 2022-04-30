<div>
	<!-- <div > -->
	<button>
		<ion-icon name="add-circle-outline" class="plus" />
	</button>

	<!-- </div> -->
</div>

<style>
	div {
		background-color: white;
		border-radius: var(--border-radius);
		outline: none;
		padding: var(--input-padding);
		height: 100px;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	button {
		background-color: inherit;
		border: none;
	}

	.plus {
		font-size: 48px;
		transition: all 300ms;
	}

	.plus:hover {
		color: var(--accent-foreground);
		transform: rotate(135deg) scale(1.4);
	}

	.plus:hover::after {
		content: 'Upload Image';
	}
</style>
