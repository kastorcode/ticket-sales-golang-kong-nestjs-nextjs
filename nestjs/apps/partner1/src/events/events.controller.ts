import { Body, Controller, Delete, Get, HttpCode, Param, Patch, Post, UseGuards } from '@nestjs/common'
import { AuthGuard } from '@app/core/auth/auth.guard'
import { EventsService } from 'libs/core/src/events/events.service'
import { CreateEventRequest } from './request/create-event.request'
import { ReserveSpotsRequest } from './request/reserve-spots.request'
import { UpdateEventRequest } from './request/update-event.request'
import { ReserveSpotsResponse } from './response/reserve-spots-response'

@Controller('events')
export class EventsController {

  constructor (private readonly eventsService : EventsService) {}

  @Post()
  create (@Body() data : CreateEventRequest) {
    return this.eventsService.create(data)
  }

  @Get()
  findAll () {
    return this.eventsService.findAll()
  }

  @Get(':id')
  findOne (@Param('id') id : string) {
    return this.eventsService.findOne(id)
  }

  @Patch(':id')
  update (@Param('id') id : string, @Body() data : UpdateEventRequest) {
    return this.eventsService.update(id, data)
  }

  @HttpCode(204)
  @Delete(':id')
  remove (@Param('id') id : string) {
    return this.eventsService.remove(id)
  }

  @UseGuards(AuthGuard)
  @Post(':id/reserve')
  async reserveSpots (@Param('id') id : string, @Body() data : ReserveSpotsRequest) {
    const tickets = await this.eventsService.reserveSpots(id, data)
    return new ReserveSpotsResponse(tickets)
  }

}
