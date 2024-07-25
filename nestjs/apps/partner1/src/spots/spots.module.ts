import { Module } from '@nestjs/common'
import { SpotsCoreModule } from 'libs/core/src/spots/spots-core.module'
import { SpotsController } from './spots.controller'

@Module({
  imports: [SpotsCoreModule],
  controllers: [SpotsController]
})
export class SpotsModule {}
