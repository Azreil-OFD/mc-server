
export default defineNuxtRouteMiddleware(async (to, from) => {
    const token = useCookie('token')
    if(!to.path) {
        return
    }

    if (to.path == '/login') {
        if (token.value !== undefined) {
            const result = await useFetch('/api/authMiddleware', {
                method: 'post',
                body: {
                    token: token.value
                }
            })
            if (result.data.value.status) {
                return navigateTo('/')
            }
        }
        return
    }

    const ignore = ['/api/getOnline','/api/authMiddleware', '/_nuxt']
    if ( ignore.includes(to.path)) {
        return
    }

    if (token.value === undefined) {
        return navigateTo('/login')
    }

    const result = await useFetch('/api/authMiddleware', {
        method: 'post',
        body: {
            token: token.value
        }
    })
    if (result.data.value.status) {
        return
    }

    return navigateTo('/login')
})