import { Test, TestingModule } from '@nestjs/testing'
import { EventsService } from 'libs/core/src/events/events.service'
import { EventsController } from './events.controller'

describe('EventsController', () => {
  let controller: EventsController

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [EventsController],
      providers: [EventsService],
    }).compile()

    controller = module.get<EventsController>(EventsController)
  })

  it('should be defined', () => {
    expect(controller).toBeDefined()
  })
})
