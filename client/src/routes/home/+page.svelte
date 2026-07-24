<script lang="ts">
    import type { PageProps } from './$types';
    import { dataItem, trackItem } from '$lib/types'
    import { getUserContext } from '$lib/context';
    import CurrentTrackCard from '$lib/components/activity/CurrentTrackCard.svelte';
    import RecentTracks from '$lib/components/activity/RecentTracks.svelte';
    import { BarChart } from 'layerchart';
	import { nonpassive } from 'svelte/legacy';

    const { data }: PageProps = $props();
    const profileData = getUserContext();

    const topTrackChartData: dataItem[] = data.topTracksData?.map((track: trackItem) => {
        return new dataItem(track.name, track.timesListened);
    });


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

        <CurrentTrackCard track={data.currentTrackData} />
    </div>

    <div class="w-full flex flex-row gap-2">
    
        <BarChart
            data={topTrackChartData}
            orientation="horizontal"
            x="value"
            y="key"
            axis="y"
            rule={false}
            padding={{ left: 4, bottom: 20, top: 20, right: 4 }}
            height={500}
            labels={true}
            props={{
                labels: {
                    textAnchor: 'end',
                    fill: 'white',
                },
                yAxis: {
                    tickLabelProps: {
                        textAnchor: "start",
                        dx: 10,
                        class: "text-sm text-taupe-200",
                    }
                },
                bars: {
                    stroke: 'none',
                    height: 50,
                }
            }}
        />
    
    </div>

    
    <div class="w-full flex flex-row gap-2">
    
        <div class="w-full">
            <RecentTracks tracks={data.recentTracksData}/>
        </div>
    
    </div>
</div>