<script lang="ts">
    import type { PageProps } from './$types';
    import { getUserContext } from '$lib/context';
    import CurrentTrackCard from '$lib/components/activity/CurrentTrackCard.svelte';
    import RecentTracks from '$lib/components/activity/RecentTracks.svelte';


    const { data }: PageProps = $props();
    const profileData = getUserContext();
</script>


<div class="w-3/5 min-h-screen bg-taupe-700 p-2 flex flex-col gap-2 rounded-md">
    <div class="w-full flex flex-row items-center justify-between gap-2">

        <div class="w-full flex flex-row gap-2">
            <img class="w-24 h-24 rounded-full bg-black" src="" alt="pfp"/>

            <div class="flex flex-col gap-1 items-center justify-center">
                <p class="text-xl text-taupe-500">Welcome back,</p>
                <p class="text-3xl font-bold text-taupe-200">{profileData ? profileData.username : "username"}</p>
            </div>
        </div>

    {#if data.currentTrackData}
        <CurrentTrackCard track={data.currentTrackData} />
    {/if}
    
    </div>

    <div class="w-full flex flex-row gap-2">
    
    <div class="w-full">
    {#if data.currentTrackData}
        <RecentTracks tracks={[data.currentTrackData, data.currentTrackData, data.currentTrackData, ]}/>
    {/if}
    </div>
    
    </div>
</div>