<script lang="ts">
    import type { navButton } from '$lib/types'
    import NavButton from './NavButton.svelte'

    const { buttons }: { buttons: navButton[] } = $props();
  
</script>

<div class="w-full border-b-purple-400 border-b-1 backdrop-blur-sm bg-purple-950/70 sticky top-0 flex flex-row gap-4 items-center justify-between pt-2 pb-2 pr-4 pl-4 z-10">

    <p class="text-taupe-200 text-2xl font-semibold">
        <a href="/">symbalx</a>
    </p>

    <div class="flex flex-row gap-2">
        {#each buttons as button}
            <NavButton text={button.text} redirect={button.redirect} image={button.image}/>
        {/each}
    </div>

</div>
