import { Body, Controller, HttpCode, Post } from '@nestjs/common';
import { PolicyIssuedEvent } from '../../../libs/contracts/src/policy-issued.event';

@Controller('notifications')
export class NotificationsController {
  @Post()
  @HttpCode(202)
  policyIssued(@Body() event: PolicyIssuedEvent) {
    console.log(`[notificación] póliza emitida ${event.number}`);
    return { accepted: true };
  }
}
