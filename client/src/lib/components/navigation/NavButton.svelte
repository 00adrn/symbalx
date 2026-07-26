<script lang="ts">
    import type { navButton } from '$lib/types'
    import { Button } from '$lib/components/ui/button/index.js'

    const { text, redirect, image }: navButton = $props();
</script>

<Button 
    variant="outline" 
    size="icon" 
    aria-label={text} 
    href={redirect}
    class="h-8 w-16 text-taupe-200 rounded-md bg-transparent hover:bg-taupe-500 hover:text-taupe-300 transition-all font-bold">
    {text}
</Button>

<!-- <div class="h-12 w-12 static top-0 rounded-full bg-taupe-600 hover:bg-taupe-500 transition-all flex items-center justify-center">
    <a aria-label={text} href={redirect} class="h-full w-full text-taupe-200 font-semibold text-lg flex items-center justify-center px-1 py-1">
        <img src={image} alt={text + " icon"} class="h-6 w-6 object-contain"/>
    </a>
</div> -->