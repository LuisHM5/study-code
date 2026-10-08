import { Module } from '@nestjs/common';
import { EventEmitterModule } from '@nestjs/event-emitter';
import { NotificationsModule } from './notifications/notifications.module';
import { PoliciesModule } from './policies/policies.module';

@Module({
  imports: [EventEmitterModule.forRoot(), PoliciesModule, NotificationsModule],
})
export class AppModule {}
