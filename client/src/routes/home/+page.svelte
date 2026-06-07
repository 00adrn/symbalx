<script lang="ts">
    import type { PageProps } from './$types';
    import CurrentTrackCard from '$lib/components/activity/CurrentTrackCard.svelte';

    const { data }: PageProps = $props();
</script>

<div class="w-full h-full bg-taupe-700 p-4 flex flex-col gap-2 rounded-md">
    <div class="w-full flex flex-row gap-2">

    {#if data.currentTrackData}
        <CurrentTrackCard track={data.currentTrackData} />
    {/if}
    
    </div>
</div>