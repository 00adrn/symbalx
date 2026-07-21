<script lang="ts">
    import type { PageProps } from './$types';
    import { getUserContext } from '$lib/context';
    import SpotifyDataCard from '$lib/components/profileInfo/SpotifyDataCard.svelte';

    const { data }: PageProps = $props();
    const profileData = getUserContext();
</script>

<div class="w-3/5 h-full bg-taupe-700 p-2 flex flex-col gap-2 rounded-md min-h-screen ">

    <div class="flex flex-row items-center gap-2">
        <div class="w-full">
            <img class="bg-black rounded-full h-24 w-24"src="" alt="pfp" />
        </div>
    </div>
    <div class="flex flex-row itms-center gap-2">
        <p class="font-bold text-taupe-200 text-4xl">{profileData ? profileData.username : "username"}</p>
    </div>
    
    <div class="flex flex-row justify-center gap-2">
        <SpotifyDataCard spotifyProfileData={data.spotifyProfileData} />
    </div>
</div>