import { login } from "~/service/login"

export default defineEventHandler(async (event) => {
  const data = await readBody(event)

  return await login(data.login , data.password)
})
