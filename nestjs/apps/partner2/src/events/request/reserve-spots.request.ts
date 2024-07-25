import { TicketKind } from '@prisma/client'

export class ReserveSpotsRequest {
  spots      : string[]
  ticketKind : TicketKind
  email      : string
}