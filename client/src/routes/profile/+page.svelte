<script lang="ts">
    import type { PageProps } from './$types';
    import { getUserContext } from '$lib/context';
    import SpotifyDataCard from '$lib/components/profileInfo/SpotifyDataCard.svelte';
    import { onMount } from 'svelte'
    import { goto } from "$app/navigation"
    import * as Avatar from "$lib/components/ui/avatar/index.js"

    const { data }: PageProps = $props();
    const profileData = getUserContext();

    onMount(() => {
        if (!profileData){
            goto("/auth/login")
        }
    });
</script>

<div class="w-4/5 h-full bg-taupe-700 p-2 flex flex-col gap-2 rounded-md min-h-screen ">

    <Avatar.Root class="w-24 h-24">
        <Avatar.Image src="" alt="profile picture" />
        <Avatar.Fallback>{profileData ? profileData.username[0] : "username"}</Avatar.Fallback>
    </Avatar.Root>

    <div class="flex flex-row itms-center gap-2">
        <p class="font-bold text-taupe-200 text-4xl">{profileData ? profileData.username : "username"}</p>
    </div>
    
    <div class="flex flex-row justify-center gap-2">
        <SpotifyDataCard spotifyProfileData={data.spotifyProfileData} />
    </div>
</div>