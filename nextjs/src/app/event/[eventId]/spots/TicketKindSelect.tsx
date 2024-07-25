"use client"

import { selectTicketTypeAction } from "~/actions"

export type TicketKindSelectProps = {
  defaultValue : "full" | "half"
  price        : number
}

export function TicketKindSelect ({ defaultValue, price } : TicketKindSelectProps) {

  const formattedFullPrice = new Intl.NumberFormat("pt-BR", {
    style: "currency", currency: "BRL"
  }).format(price)

  const formattedHalfPrice = new Intl.NumberFormat("pt-BR", {
    style: "currency", currency: "BRL"
  }).format(price / 2)

  return (
    <>
      <label htmlFor="ticket-type">Choose ticket type</label>
      <select
        name="ticket-type"
        id="ticket-type"
        className="mt-2 rounded-lg bg-input px-4 py-[14px]"
        defaultValue={defaultValue}
        onChange={async e => {
          await selectTicketTypeAction(e.target.value as TicketKindSelectProps["defaultValue"])
        }}
      >
        <option value="full">Full - {formattedFullPrice}</option>
        <option value="half">Half - {formattedHalfPrice}</option>
      </select>
    </>
  )

}