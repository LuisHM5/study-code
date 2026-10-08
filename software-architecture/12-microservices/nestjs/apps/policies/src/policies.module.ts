import { Module } from '@nestjs/common';
import { NotificationsClient } from './notifications.client';
import { PoliciesController } from './policies.controller';
import { PoliciesService } from './policies.service';

@Module({
  controllers: [PoliciesController],
  providers: [PoliciesService, NotificationsClient],
})
export class PoliciesModule {}
