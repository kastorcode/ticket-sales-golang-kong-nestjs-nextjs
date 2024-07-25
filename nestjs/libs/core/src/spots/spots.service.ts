import { Injectable } from '@nestjs/common'
import { SpotStatus } from '@prisma/client'
import { PrismaService } from 'libs/core/src/prisma/prisma.service'
import { CreateSpotDto } from './dto/create-spot.dto'
import { UpdateSpotDto } from './dto/update-spot.dto'

@Injectable()
export class SpotsService {

  constructor (private prismaService : PrismaService) {}

  async create (eventId : string, createSpotDto : CreateSpotDto) {
    const event = await this.prismaService.event.findUnique({ where: { id: eventId }})
    if (!event) throw new Error('Event not found')
    return this.prismaService.spot.create({ data: {
      ...createSpotDto, eventId, status: SpotStatus.available
    }})
  }

  findAll (eventId : string) {
    return this.prismaService.spot.findMany({ where: { eventId } })
  }

  findOne (eventId : string, spotId : string) {
    return this.prismaService.spot.findUnique({ where: { id: spotId, eventId } })
  }

  update (eventId : string, spotId : string, data : UpdateSpotDto) {
    return this.prismaService.spot.update({ where: { id: spotId, eventId }, data })
  }

  remove (eventId : string, spotId : string) {
    return this.prismaService.spot.delete({ where: { id: spotId, eventId } })
  }

}
