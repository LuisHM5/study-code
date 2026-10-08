import { Module } from '@nestjs/common';
import { PolicyController } from './controllers/policy.controller';
import { PolicyRepository } from './repositories/policy.repository';
import { MailService } from './services/mail.service';
import { PolicyService } from './services/policy.service';

// Un solo módulo: las carpetas se organizan por tipo técnico, no por feature.
@Module({
  controllers: [PolicyController],
  providers: [PolicyService, MailService, PolicyRepository],
})
export class AppModule {}
