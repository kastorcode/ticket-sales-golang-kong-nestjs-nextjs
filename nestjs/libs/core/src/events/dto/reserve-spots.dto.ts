import { TicketKind } from '@prisma/client'

export class ReserveSpotsDto {
  spots      : string[]
  ticketKind : TicketKind
  email      : string
}