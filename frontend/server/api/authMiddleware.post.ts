import { authMiddleware } from "~/service/auth"


export default defineEventHandler(async (event) => {
  const data = await readBody(event)
  

  const result = await authMiddleware(data.token)
  return result
})
