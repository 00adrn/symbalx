<script lang="ts">
    import * as Item from "$lib/components/ui/item/index.js"
    import { trackItem } from '$lib/types'

    const { tracks }: { tracks: trackItem[] } = $props();
</script>

{#if tracks}
<div class="w-full flex flex-row items-center gap-8">

    <div class="flex flex-col w-full">
        <!-- <p class="text-xl text-taupe-500">Recent tracks:</p> -->
        {#each tracks as track, index}
        <Item.Root >
            <Item.Media variant="image" class="h-10 w-10">
                <img alt="Current Track" src={track.getImage} />
            </Item.Media>
            <Item.Content>
                <Item.Title class="font-semibold text-md text-violet-950">{track.name}</Item.Title>
                <Item.Description class="text-sm text-purple-500">{track.getArtistString}</Item.Description>
            </Item.Content>
        </Item.Root>
        <div class="border-b-1 w-full border-gray-400"></div>
        {/each}
    </div>
</div> 
{/if}