<script lang="ts">
    import { goto } from '$app/navigation'
    import LineTextInput from '$lib/components/input/LineTextInput.svelte';
    import PillButton from '$lib/components/navigation/PillButton.svelte';

    let emailInput: string = $state('');
    let passwordInput: string = $state('');
    let usernameInput: string = $state('');
    let isLogin: boolean = $state(true);

</script>

<div class="w-3/5 min-h-screen bg-taupe-700 p-4 flex flex-col gap-2 rounded-md">
    <div class="w-full flex flex-col items-center justify-center gap-2">

        <div class="w-1/2 flex flex-row gap-1 justify-evenly">
            <button onclick={() => { isLogin = true; }}
                class="w-1/2 rounded-md bg-taupe-500 py-1 font-bold text-taupe-200 border-4 border-taupe-800">
                Login
            </button>
            <button onclick={() => { isLogin = false; }}
                class="w-1/2 rounded-md bg-taupe-500 py-1 font-bold text-taupe-200 border-4 border-taupe-800">
                Register
            </button>
        </div>

        {#if isLogin}
        <form class="w-1/2 flex flex-col gap-4 bg-taupe-800 p-2 rounded-md items-center justify-between text-taupe-200 text-xl
                border-md border-taupe-200" method="POST" action="?/login">

            <label class="text-taupe-200 font-semibold flex-col">
                <p>Email:</p>
                <input name="email" type="email" bind:value={emailInput}
                    class="bg-taupe-500 rounded-md px-1">
            </label>
            <label class="text-taupe-200 font-semibold flex-col">
                <p>Password:</p>
                <input name="password" type="password" bind:value={passwordInput}
                    class="bg-taupe-500 rounded-md px-1">
            </label>

            <button class="rounded-full bg-taupe-500 font-semibold px-16 py-1">Log in</button>
        </form>
        {:else}
        <form class="w-1/2 flex flex-col gap-4 bg-taupe-800 p-2 rounded-md items-center justify-center text-taupe-200 text-xl
                border-md border-taupe-200" method="POST" action="?/register">
            
            <label class="text-taupe-200 font-semibold flex-col">
                <p>Username:</p>
                <input name="username" type="text" bind:value={usernameInput}
                    class="bg-taupe-500 rounded-md px-1">
            </label>
            <label class="text-taupe-200 font-semibold flex-col">
                <p>Email:</p>
                <input name="email" type="email" bind:value={emailInput}
                    class="bg-taupe-500 rounded-md px-1">
            </label>
            <label class="text-taupe-200 font-semibold flex-col">
                <p>Password:</p>
                <input name="password" type="password" bind:value={passwordInput}
                    class="bg-taupe-500 rounded-md px-1">
            </label>

            <button class="rounded-full bg-taupe-500 font-semibold px-16 py-1">Register</button>
        </form>
        {/if}
    </div>
</div>