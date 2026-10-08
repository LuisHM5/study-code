import { Module } from '@nestjs/common';
import { PolicyIssuedListener } from './listeners/policy-issued.listener';

@Module({ providers: [PolicyIssuedListener] })
export class NotificationsModule {}
