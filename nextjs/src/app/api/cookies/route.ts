import { cookies } from "next/headers"

export function DELETE () {
  const cookieStore = cookies()
  const allCookies = cookieStore.getAll()
  allCookies.forEach(cookie => cookieStore.delete(cookie.name))
  return new Response(null, { status: 204 })
}