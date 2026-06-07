<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/navigation/Navbar.svelte'

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="bg-taupe-800">
	<Navbar />
	<div class="min-h-screen w-full py-2 px-4 flex flex-col gap-2">
		{@render children()}
	</div>
</div>