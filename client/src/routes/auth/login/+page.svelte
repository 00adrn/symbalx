<script lang="ts">
    import { Button } from "$lib/components/ui/button/index.js"
    import { Input } from "$lib/components/ui/input/index.js"
    import { Label } from "$lib/components/ui/label/index.js"
    import * as Card from "$lib/components/ui/card/index.js"


    let emailInput: string = $state('');
    let passwordInput: string = $state('');
    let usernameInput: string = $state('');
    let isLogin: boolean = $state(true);

</script>

<div class="w-full min-h-screen bg-gray-200 flex items-center justify-center p-4">
	{#if isLogin}
		<Card.Root class="w-full max-w-sm">
			<Card.Header>
				<Card.Title class="text-2xl">Login</Card.Title>
				<Card.Description>Enter your email below to login to your account.</Card.Description>
			</Card.Header>
			<form method="POST" action="?/login">
				<Card.Content class="grid gap-4">
					<div class="grid gap-2">
						<Label for="email">Email</Label>
						<Input id="email" type="email" name="email" placeholder="m@example.com" required bind:value={emailInput} />
					</div>
					<div class="grid gap-2">
						<Label for="password">Password</Label>
						<Input id="password" type="password" name="password" required bind:value={passwordInput} />
					</div>
				</Card.Content>
				<Card.Footer class="flex flex-col gap-4">
					<Button type="submit" class="w-full">Sign in</Button>
					<div class="mt-4 text-center text-sm">
						Don't have an account?
						<button type="button" onclick={() => isLogin = false} class="underline">
							Sign up
						</button>
					</div>
				</Card.Footer>
			</form>
		</Card.Root>
	{:else}
		<Card.Root class="w-full max-w-sm">
			<Card.Header>
				<Card.Title class="text-2xl">Sign Up</Card.Title>
				<Card.Description>Enter your information to create an account.</Card.Description>
			</Card.Header>
			<form method="POST" action="?/register">
				<Card.Content class="grid gap-4">
					<div class="grid gap-2">
						<Label for="username">Username</Label>
						<Input id="username" name="username" placeholder="max_robinson" required bind:value={usernameInput} />
					</div>
					<div class="grid gap-2">
						<Label for="email">Email</Label>
						<Input id="email" type="email" name="email" placeholder="m@example.com" required bind:value={emailInput} />
					</div>
					<div class="grid gap-2">
						<Label for="password">Password</Label>
						<Input id="password" type="password" name="password" required bind:value={passwordInput} />
					</div>
				</Card.Content>
				<Card.Footer class="flex flex-col gap-4">
					<Button type="submit" class="w-full">Create account</Button>
					<div class="mt-4 text-center text-sm">
						Already have an account?
						<button type="button" onclick={() => isLogin = true} class="underline">
							Sign in
						</button>
					</div>
				</Card.Footer>
			</form>
		</Card.Root>
	{/if}
</div>