<script lang="ts">
    import { handleLogin, handleSignUp} from '$lib/auth'
    import { goto } from '$app/navigation';

    import LineTextInput from '$lib/components/input/LineTextInput.svelte';
    import PillButton from '$lib/components/navigation/PillButton.svelte';

    let emailInput: string = $state('');
    let passwordInput: string = $state('');
    let usernameInput: string = $state('');
    let isLogin: boolean = $state(true);

    const switchType = () => {
        usernameInput = "";
        emailInput = "";
        passwordInput = "";
        isLogin = !isLogin;
    }

    const onSubmit = async () => {
        if (emailInput === "" || passwordInput === "" || (!isLogin && usernameInput === "")) {
            return;
        }

        const res = isLogin ? await handleLogin(emailInput, passwordInput) : await handleSignUp(usernameInput, emailInput, passwordInput);

        if (res) {
            await goto('/home', { replaceState: true });
        }
    }
</script>

<div class="w-3/5 min-h-screen bg-taupe-700 p-4 flex flex-col gap-2 rounded-md">
    <div class="w-full flex flex-row items-center justify-center gap-2">

        <div class="w-1/2 flex flex-col gap-4 bg-taupe-800 p-2 rounded-md items-center justify-center text-taupe-200 text-xl
                border-md border-taupe-200">
            <div class="w-full justify-between flex flex-row">

                <p class="text-3xl text-taupe-200 font-bold">
                    {isLogin ? "Login" : "Sign Up"}
                </p>

                <div class="min-w-1/3 max-w-1/2">
                    <PillButton text={isLogin ? "I'm new here" : "I have an account"} 
                        onClick={switchType}/>
                </div>

            </div>

            {#if !isLogin}
                <LineTextInput bind:textInput={usernameInput} placeholder="Username" />
            {/if}
            <LineTextInput bind:textInput={emailInput} placeholder="Email" />
            <LineTextInput bind:textInput={passwordInput} placeholder="Password" />
            <div class="w-2/3 px-2">
                <PillButton text="Submit" onClick={onSubmit} />
            </div>

        </div>
    </div>
</div>