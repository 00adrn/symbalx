<script lang="ts">
    import { trackItem } from '$lib/types'
    import * as Item from "$lib/components/ui/item/index.js"

    const { track }: { track: trackItem | null } = $props();

    const genArtistString = () : string => {
        let artistString = '';

        if (!track) return artistString;

        track.artists.forEach((artist, index) => {
            artistString += artist.name;
            if (index != track.artists.length -1) artistString += ", ";
        })

        return artistString;
    }
    //<img class="h-18 w-18 rounded-md bg-transparent" alt="Current Track" src={track.getImage} />
</script>

<div class="w-full min-h-full bg-transparent flex flex-row items-center gap-8 rounded-md px-2 py-2">
    {#if track}
    <Item.Root >
        <Item.Media variant="image" class="h-16 w-16">
            <img alt="Current Track" src={track.getImage} />
        </Item.Media>
        <Item.Content>
            <Item.Title class="font-semibold text-md text-violet-950"><span class="text-sm text-gray-400">Currently Playing - </span>{track.name}</Item.Title>
            <Item.Description class="text-sm text-purple-400">{track.getArtistString}</Item.Description>
        </Item.Content>
    </Item.Root>
    {:else}
    <Item.Root >
        <Item.Media variant="image" class="h-16 w-16">
            <img alt="Current Track" src=""/>
        </Item.Media>
        <Item.Content>
            <Item.Title class="font-semibold text-md text-violet-950"><span class="text-sm text-gray-400">No track playing...</span></Item.Title>
            <Item.Description class="text-sm text-purple-400">-</Item.Description>
        </Item.Content>
    </Item.Root>
    {/if}
</div> 