<script lang="ts">
    import type { PageProps } from './$types';
    import { getUserContext } from '$lib/context';
    import SpotifyDataCard from '$lib/components/profileInfo/SpotifyDataCard.svelte';
    import * as Avatar from "$lib/components/ui/avatar/index.js"

    const { data }: PageProps = $props();
    const profileData = getUserContext();

</script>

<div class="w-full md:w-full min-h-screen bg-gray-200 px-4 py-4 md:px-64 flex flex-col gap-2">

    <Avatar.Root class="w-24 h-24">
        <Avatar.Image src="" alt="profile picture" />
        <Avatar.Fallback>{profileData ? profileData.username[0] : "username"}</Avatar.Fallback>
    </Avatar.Root>

    <div class="flex flex-row items-center gap-2">
        <p class="font-bold text-purple-900 text-3xl sm:text-4xl">{profileData ? profileData.username : "username"}</p>
    </div>
    
    <div class="flex flex-row justify-center gap-2">
        <SpotifyDataCard spotifyProfileData={data.spotifyProfileData} />
    </div>
</div>