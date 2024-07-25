import { NestFactory } from '@nestjs/core'
import { EventsService } from '@app/core/events/events.service'
import { PrismaService } from '@app/core/prisma/prisma.service'
import { SpotsService } from '@app/core/spots/spots.service'
import { Partner1Module } from '../../../apps/partner1/src/partner1.module'
import { Partner2Module } from '../../../apps/partner2/src/partner2.module'

interface Event {
  id           : string
  name         : string
  location     : string
  organization : string
  rating       : string
  date         : string
  imageUrl     : string
  capacity     : number
  price        : number
  partnerId    : number
}

interface Spot {
  id       : string
  name     : string
  eventId  : string
  reserved : boolean
  status   : string
  ticketId : string
}

async function fixture (databaseUrl : string, module : Partner1Module|Partner2Module, events : Event[]) {

  process.env.DATABASE_URL=databaseUrl

  const app = await NestFactory.createApplicationContext(module)

  const prismaService = app.get<PrismaService>(PrismaService)
  await prismaService.reservationHistory.deleteMany({})
  await prismaService.ticket.deleteMany({})
  await prismaService.spot.deleteMany({})
  await prismaService.event.deleteMany({})

  const eventsService = app.get(EventsService)

  const createdEvents = await Promise.all(
    events.map(async event => {
      await eventsService.create({
        id         : event.id,
        name       : event.name,
        description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam at purus at lacus rutrum pellentesque et in tellus. Sed egestas nunc in purus vestibulum lacinia.',
        date       : event.date,
        price      : event.price
      })
      return event.id
    })
  )

  const spotsService = app.get(SpotsService)

  await Promise.all(createdEvents.map(async eventId => {
    const response = await fetch(`http://host.docker.internal:8080/events/${eventId}/spots`)
    const {spots} : {event:Event,spots:Spot[]} = await response.json()
    await Promise.all(spots.map(async spot => {
      await spotsService.create(eventId, {id:spot.id,name:spot.name})
    }))
  }))

  await app.close()

}

async function bootstrap () {

  const response = await fetch('http://host.docker.internal:8080/events')
  const {events} : {events:Event[]} = await response.json()

  const partner1Events : Event[] = []
  const partner2Events : Event[] = []
  events.forEach(event => {
    if      (event.partnerId === 1) partner1Events.push(event)
    else if (event.partnerId === 2) partner2Events.push(event)
  })

  await fixture(
    "postgresql://root:root@host.docker.internal:5433/partner1", Partner1Module, partner1Events
  )
  await fixture(
    "postgresql://root:root@host.docker.internal:5433/partner2", Partner2Module, partner2Events
  )

}

bootstrap()