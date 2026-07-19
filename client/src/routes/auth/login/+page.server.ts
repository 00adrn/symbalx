import type { Actions } from './$types'
import { redirect } from '@sveltejs/kit'

export const actions = {
    login: async ({ cookies, request, fetch }) => {
        const data = await request.formData();
        const password = data.get("password");
        const email = data.get("email")

        const resp = await fetch('/api/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email, password })
        });

        if (resp.ok) {
            throw redirect(303, '/home');
        }

        return {
            success: false,
        }
    },
    register: async ({ cookies, request, fetch}) => {
        const data = await request.formData();
        const password = data.get("password");
        const email = data.get("email")
        const username = data.get("username")

        const resp = await fetch('/api/auth/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, email, password })
        });

        if (resp.ok) {
            throw redirect(303, '/home');
        }

        return {
            success: false,
        }
    }
} satisfies Actions;