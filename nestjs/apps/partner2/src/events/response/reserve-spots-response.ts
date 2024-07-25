import { Spot, Ticket } from '@prisma/client'

type TicketWithSpot = Ticket & { Spot : Spot }

export class ReserveSpotsResponse {

  constructor(readonly tickets : TicketWithSpot[]) {}

  toJSON () {
    return this.tickets.map(ticket => ({
      id: ticket.id,
      email: ticket.email,
      spot: ticket.Spot.name,
      ticketKind: ticket.ticketKind,
      status: ticket.Spot.status,
      eventId: ticket.Spot.eventId
    }))
  }

}