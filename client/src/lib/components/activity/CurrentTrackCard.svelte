<script lang="ts">
    import { trackItem } from '$lib/types'

    const { track }: { track: trackItem } = $props();

    const genArtistString = () => {
        let artistString = '';

        track.artists.forEach((artist, index) => {
            artistString += artist.name;
            if (index != track.artists.length -1) artistString += ", ";
        })

        return artistString;
    }
</script>

<div class="w-full min-h-full bg-taupe-800 p-2 flex flex-row items-center gap-8 rounded-md pr-2 pl-2 pt-2 pb-2">

    <img class="h-20 w-20 rounded-md" alt="Current Track" src={track.getImage} />

    <div class="flex flex-col gap-1 pt-1 pb-1">
        <p class="font-bold text-sm text-taupe-500">Currently listening to:</p>
        <p class="font-semibold text-xl text-taupe-200">{track.name}</p>
        <p class="font-bold text-sm text-taupe-500">{genArtistString()}</p>
    </div>

</div> 