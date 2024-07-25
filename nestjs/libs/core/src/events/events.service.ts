import { Injectable } from '@nestjs/common'
import { Prisma, SpotStatus, TicketStatus } from '@prisma/client'
import { PrismaService } from 'libs/core/src/prisma/prisma.service'
import { CreateEventDto } from './dto/create-event.dto'
import { ReserveSpotsDto } from './dto/reserve-spots.dto'
import { UpdateEventDto } from './dto/update-event.dto'

@Injectable()
export class EventsService {

  constructor (private prismaService : PrismaService) {}

  create (data : CreateEventDto) {
    return this.prismaService.event.create({
      data: { ...data, date: new Date(data.date) }
    })
  }

  findAll () {
    return this.prismaService.event.findMany()
  }

  findOne (id : string) {
    return this.prismaService.event.findUnique({ where: { id } })
  }

  update (id : string, data : UpdateEventDto) {
    return this.prismaService.event.update({
      data,
      where: { id }
    })
  }

  remove (id : string) {
    return this.prismaService.event.delete({ where: { id } })
  }

  async reserveSpots (eventId : string, dto : ReserveSpotsDto) {
    const spots = await this.prismaService.spot.findMany({
      where: { eventId, name: { in: dto.spots }}
    })
    if (spots.length !== dto.spots.length) {
      const foundSpotsNames = spots.map(({ name }) => name)
      const notFoundSpots = dto.spots.filter(name => !foundSpotsNames.includes(name))
      throw new Error(`Spots ${notFoundSpots.join(', ')} not found`)
    }
    try {
      return await this.prismaService.$transaction(async prisma => {
        await prisma.reservationHistory.createMany({
          data: spots.map(spot => ({
            spotId: spot.id,
            ticketKind: dto.ticketKind,
            email: dto.email,
            status: TicketStatus.reserved
          }))
        })
        await prisma.spot.updateMany({
          where: {
            id: {
              in: spots.map(({ id }) => id)
            }
          },
          data: {
            status: SpotStatus.reserved
          }
        })
        const tickets = await Promise.all(spots.map(({ id }) => prisma.ticket.create({
          data: {
            spotId: id,
            ticketKind: dto.ticketKind,
            email: dto.email
          },
          include: { Spot: true }
        })))
        return tickets
      }, { isolationLevel: Prisma.TransactionIsolationLevel.ReadCommitted })
    }
    catch (error) {
      if (error instanceof Prisma.PrismaClientKnownRequestError) {
        switch (error.code) {
          case 'P2002':
          case 'P2034': throw new Error('Some spot(s) already reserved')
        }
      }
      throw error
    }
  }

}
