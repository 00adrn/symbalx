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
    class="h-8 w-16 text-purple-200 rounded-full bg-transparent hover:bg-purple-800 hover:text-taupe-200 transition-all font-bold">
    {text}
</Button>