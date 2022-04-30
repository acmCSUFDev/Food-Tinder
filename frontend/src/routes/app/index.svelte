<script>
	import * as svelte from 'svelte';
	import * as api from '$lib/api';
	import Loading from '$lib/components/loading.svelte';
	import NavbarThing from '../../lib/components/nav.svelte';

	const fallbackAvatar = 'https://fonts.gstatic.com/s/i/materialiconsoutlined/face/v15/24px.svg';

	let loading = true;

	/** @type {Array<api.PostListing>} */
	let posts = [];

	// TODO: redesign the API lol.
	/** @type {Map<string, api.User>} */
	let users = {};

	async function loadMore() {
		let nextPosts = [];
		if (posts.length > 0) {
			nextPosts = await api.getNextPosts(posts[posts.length - 1].id);
		} else {
			nextPosts = await api.getNextPosts();
		}

		posts.push(...nextPosts);

		for (const post of nextPosts) {
			if (users[post.username]) {
				continue;
			}

			try {
				users[post.username] = await api.getUser(post.username);
			} catch (err) {
				// Just ignore this user.
				users[post.username] = {};
			}
		}
	}

	function likePost(post) {
		const liked = !post.liked;
		api.likePost(post.id, { like: liked }).then(() => {
			post.liked = liked;
			post.likes = post.likes + (liked ? +1 : -1);
			posts = posts; // force Svelte to reload.
		});
	}

	svelte.onMount(() => {
		loadMore().then(() => (loading = false));
	});
</script>

<main class="background">
	<NavbarThing />
	{#if loading}
		<Loading />
	{:else}
		<div class="posts no-scrollbar">
			{#each posts as post}
				<section id="post-{post.id}" class="no-scrollbar">
					<div class="post-content">
						<img class="post-cover" alt="" src={api.assetURL(post.images[0]) || ''} />
						<div class="post-body">
							<div class="post-user">
								<img
									class="avatar"
									alt={users[post.username].display_name || post.username}
									src={api.assetURL(users[post.username].avatar) || fallbackAvatar}
								/>
								<p class="names">
									{#if users[post.username].display_name}
										<span class="display_name">{users[post.username].display_name}</span>
									{/if}
									<span class="username">{post.username}</span>
								</p>
								<button
									title="Like this post"
									class="like"
									class:liked={post.liked}
									on:click={() => likePost(post)}
								>
									<div class="like-icon" />
									<span class="like-count">{post.likes}</span>
								</button>
							</div>
							<p class="post-description">{post.description}</p>
							<div class="post-images">
								{#each post.images.slice(1) as asset}
									<img src={api.assetURL(asset)} alt="" />
								{/each}
							</div>
						</div>
					</div>
				</section>
			{/each}
		</div>
	{/if}
</main>

<style>
	main {
		height: 100%;
	}

	.posts {
		display: grid;
		gap: 12px;
		grid-auto-flow: column;
		grid-template-columns: max-content;

		height: 100%;

		overflow-x: scroll;
		overflow-y: hidden;
		scroll-snap-type: x mandatory;
	}

	.posts > * {
		--width: min(400px, calc(100vw - 16px));

		width: var(--width);
		scroll-snap-align: center;
		overflow-y: auto;

		/* Use padding to make it appear like the whole card shifts up, not like
		 * it's in a bounded box. */
		padding-top: calc(50vh - (500px / 2));
		padding-bottom: 25px;

		/* Workaround for shadows not working due to overflow-y: auto. */
		padding-left: 15px;
		padding-right: 15px;
		margin: 0 -15px;
	}

	/* Hacks to pad just enough space to be able to center the first and last
	 * children. */
	.posts > *:first-child {
		margin-left: calc(50vw - (var(--width) / 2));
	}
	.posts > *:last-child {
		margin-right: calc(50vw - (var(--width) / 2));
	}

	.post-content {
		box-shadow: 0 0 10px -6px rgba(0, 0, 0, 0.55), 0 5px 20px -8px rgba(0, 0, 0, 0.35);
		overflow: hidden;
		border-radius: var(--border-radius);

		/* This magically squishes the image and the body together, while block
		 * doesn't, and flex introduces weird inconsistency. */
		display: grid;
		grid-auto-rows: max-content;
	}

	.post-content > img.post-cover {
		object-fit: cover;
		height: 500px;
		width: 100%;
	}

	.post-user {
		display: flex;
		align-items: center;
	}

	.post-user .avatar {
		width: 38px;
		height: 38px;
		object-fit: cover;
		border-radius: 99px;
		margin-right: 0.5em;
		box-shadow: 0 2px 6px -2px rgba(0, 0, 0, 0.85);
	}

	.post-user p.names {
		display: flex;
		flex-direction: column;
		flex: 1;
	}

	.post-user p.names .display_name + .username {
		font-size: 0.8em;
		opacity: 0.8;
		line-height: 0.8em;
	}

	.like {
		border-radius: 99px;
		border: 2px solid black;
		outline: none;
		background: none;
		opacity: 0.25;
		display: flex;
		width: 38px;
		height: 38px;
		transition: linear 100ms all;
		position: relative;
	}

	.like.liked {
		--color: #ff0033;
		opacity: 1;
		border-color: var(--color);
	}

	.like:hover {
		opacity: 0.65;
	}

	.like > .like-count {
		position: absolute;
		bottom: -2px;
		right: -2px;
		background: white;
		width: 1.1em;
		height: 1.1em;
		border-radius: 99px;
		line-height: 1em;
	}

	.like > .like-icon {
		mask: url(https://fonts.gstatic.com/s/i/materialicons/favorite/v17/24px.svg) no-repeat center;
		background-color: black;
		width: 100%;
		height: 100%;
	}

	.like.liked > .like-icon {
		background-color: var(--color);
	}

	.post-body {
		background: white;
		padding: 22px;
	}

	.post-body > *:not(:last-child):not(:empty) {
		margin-bottom: 16px;
	}

	.post-body p {
		margin: 0;
		padding: 0;
	}

	.post-body p.post-description {
		white-space: break-spaces;
	}

	.post-images {
		display: grid;
		gap: 12px;
		grid-template-rows: max-content;
		grid-template-columns: repeat(3, 1fr); /* TODO: make adaptive */
	}

	.post-images > img {
		width: 100%;
		height: 100%;
		aspect-ratio: 1;
		object-fit: cover;
	}
</style>
