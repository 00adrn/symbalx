<script lang="ts">
    import { trackItem } from '$lib/types'

    const { tracks }: { tracks: trackItem[] } = $props();

    const genArtistString = (track: trackItem) => {
        let artistString = '';

        track.artists.forEach((artist, index) => {
            artistString += artist.name;
            if (index != track.artists.length -1) artistString += ", ";
        })

        return artistString;
    }
</script>

{#if tracks}
<div class="w-full bg-taupe-800 flex flex-row items-center gap-8 rounded-md px-1 py-.5">

    <div class="flex flex-col gap-2 px-1 py-2 w-full">
        <!-- <p class="text-xl text-taupe-500">Recent tracks:</p> -->
        {#each tracks as track, index}
        <div class="w-full flex flex-row items-center justify-between">
            <img class="h-16 w-16 rounded-md" alt="Current Track" src={track.getImage} />
            <p class="font-semibold text-md text-taupe-200">{track.name}</p>
            <p class="font-semibold text-md text-taupe-500">{genArtistString(track)}</p>
        </div>
        {/each}
    </div>
</div> 
{/if}