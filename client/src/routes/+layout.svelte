<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/navigation/Navbar.svelte'
	import { setUserContext } from '$lib/context'
	import type { LayoutProps } from './$types'

	let { data, children }: LayoutProps = $props();
	setUserContext(data.profileData)
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="bg-taupe-800 min-h-screen">
	<Navbar />
	<div class="min-h-screen justify-center items-center py-2 px-4 flex flex-col gap-2">
		{@render children()}
	</div>
</div>