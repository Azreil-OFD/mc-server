export default defineEventHandler(async (event) => {

  return await fetch('http://localhost:8281/mcrest/player')
})
