import { Module } from '@nestjs/common';
import { PolicyIssuedListener } from './policy-issued.listener';

@Module({ providers: [PolicyIssuedListener] })
export class NotificationsModule {}
