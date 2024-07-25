"use server"

import { revalidateTag } from "next/cache"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

export async function selectSpotAction (eventId : string, spotName : string) {
  const cookieStore = cookies()
  const spots = JSON.parse(cookieStore.get("spots")?.value || "[]")
  spots.push(spotName)
  const uniqueSpots = Array.from(new Set(spots))
  cookieStore.set("spots", JSON.stringify(uniqueSpots))
  cookieStore.set("eventId", eventId)
}

export async function unSelectSpotAction (spotName : string) {
  const cookieStore = cookies()
  const spots = JSON.parse(cookieStore.get("spots")?.value || "[]")
  const newSpots = spots.filter((spot : string) => spot !== spotName)
  cookieStore.set("spots", JSON.stringify(newSpots))
}

export async function clearSpotsAction () {
  const cookieStore = cookies()
  cookieStore.delete("ticketKind")
  cookieStore.delete("spots")
  cookieStore.delete("eventId")
}

export async function selectTicketTypeAction (ticketKind : "full" | "half") {
  const cookieStore = cookies()
  cookieStore.set("ticketKind", ticketKind)
}

export async function checkoutAction (prevState : any, { cardHash, email } : { cardHash : string, email : string }) {
  const cookieStore = cookies()
  const eventId = cookieStore.get("eventId")?.value
  const spots = JSON.parse(cookieStore.get("spots")?.value || "[]")
  const ticketKind = cookieStore.get("ticketKind")?.value || "full"
  const response = await fetch(`${process.env.GOLANG_API_URL}/checkout`, {
    method: "POST",
    headers: {
      "apikey": process.env.GOLANG_API_TOKEN as string,
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ eventId, cardHash, email, spots, ticketKind })
  })
  if (!response.ok) return { error: response.statusText }
  revalidateTag(`events/${eventId}`)
  redirect(`/checkout/${eventId}/success`)
}