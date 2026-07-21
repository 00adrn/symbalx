<script lang="ts">
    import { trackItem } from '$lib/types'

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

<div class="w-full min-h-full bg-taupe-800 flex flex-row items-center gap-8 rounded-md px-2 py-2">
    {#if track}
    <div class="flex flex-col gap-1">
        <p class="font-bold text-sm text-taupe-500">Currently listening to:</p>
        <p class="font-semibold text-xl text-taupe-200">{track.name}</p>
        <p class="font-bold text-sm text-taupe-500">{genArtistString()}</p>
    </div>
    {:else}
    <div class="flex flex-col gap-1">
        <p class="font-bold text-sm text-taupe-500">-</p>
        <p class="font-semibold text-xl text-taupe-200">No track currently playing</p>
        <p class="font-bold text-sm text-taupe-500">-</p>
    </div>
    {/if}
</div> 