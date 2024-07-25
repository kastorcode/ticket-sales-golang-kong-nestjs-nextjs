import { cookies } from "next/headers"
import { DeleteAllCookies } from '~/components/DeleteAllCookies'
import { Title } from "~/components/Title"
import { EventModel } from "~/models"

type CheckoutSuccessPageProps = {
  params : { eventId : string }
}

async function getEvent (eventId : string) : Promise<EventModel> {
  const response = await fetch(`${process.env.GOLANG_API_URL}/events/${eventId}`, {
    headers: {
      "apikey": process.env.GOLANG_API_TOKEN as string
    },
    cache: "no-store",
    next: {
      tags: [`events/${eventId}`]
    }
  })
  return response.json()
}

export default async function CheckoutSuccessPage ({ params } : CheckoutSuccessPageProps) {

  const event = await getEvent(params.eventId)
  const cookieStore = cookies()
  const selectedSpots = JSON.parse(cookieStore.get("spots")?.value || "[]")

  return (
    <>
    <DeleteAllCookies />
    <main className="flex flex-col flex-wrap items-center mt-10 min-h-full">
      <Title>Purchase made successfully!</Title>
      <div className="mb-4 flex max-h-[250px] w-full max-w-[478px] flex-col gap-y-6 rounded-2xl bg-secondary p-4">
        <Title>Purchase summary</Title>
        <p className="font-semibold">
          Event {event.name}
          <br />
          Location {event.location}
          <br />
          Date {" "}
          {new Date(event.date).toLocaleDateString("pt-BR", {
            weekday: "long",
            day: "2-digit",
            month: "2-digit",
            year: "numeric"
          })}
        </p>
        <p className="font-semibold text-white">Chosen spots: {selectedSpots.join(", ")}</p>
      </div>
    </main>
    </>
  )

}