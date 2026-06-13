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

<div class="w-full bg-taupe-800 p-2 flex flex-row items-center gap-8 rounded-md px-2 py-1">


    <div class="flex flex-col gap-1 px-1 py-1 w-full">
        <p class="text-xl text-taupe-500">Recent tracks:</p>
        {#each tracks as track, index}
        <div class="w-full flex flex-row items-center justify-between">
            <img class="h-18 w-18 rounded-md" alt="Current Track" src={track.getImage} />
            <p class="font-semibold text-md text-taupe-200">{track.name}</p>
            <p class="font-semibold text-md text-taupe-500">{genArtistString(track)}</p>
        </div>
        {/each}
    </div>

</div> 