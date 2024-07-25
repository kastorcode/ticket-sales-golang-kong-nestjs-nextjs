import { CanActivate, ExecutionContext, Injectable } from '@nestjs/common'
import { ConfigService } from '@nestjs/config'
import { Request } from 'express'

@Injectable()
export class AuthGuard implements CanActivate {

  constructor (private configService : ConfigService) {}

  canActivate (context : ExecutionContext) : boolean {
    const request = context.switchToHttp().getRequest<Request>()
    const requestToken = request.headers['x-api-token']
    const envToken = this.configService.get('X_API_TOKEN')
    return requestToken === envToken
  }

}
