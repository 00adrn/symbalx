<script lang="ts">
    import type { navButton } from '$lib/types'
    import NavButton from './NavButton.svelte'

    const buttons: navButton[] = $state([
        {
            text: "Home",
            redirect: "/home"
        },
        {
            text: "Profile",
            redirect: "/profile"
        },
        {
            text: "spotifyauthlogin",
            redirect: "/auth/spotify/login"
        }
    ]);    
</script>

<div class="w-full border-b-taupe-400 border-b-1 backdrop-blur-sm bg-taupe-600/30 sticky top-0 flex flex-row gap-4 items-center justify-between pt-2 pb-2 pr-4 pl-4">

    <p class="text-taupe-200 text-2xl font-semibold">
        <a href="/">symbalx</a>
    </p>

    <div class="flex flex-row gap-2">
        {#each buttons as button}
            <NavButton text={button.text} redirect={button.redirect} />
        {/each}
    </div>

</div>

<style>

</style>