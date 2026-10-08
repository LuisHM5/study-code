import { Module } from '@nestjs/common';
import { NotificationService } from '../notification/notification.service';
import { PolicyStore } from './entities/policy.store';
import { GetPolicyController } from './get-policy/get-policy.controller';
import { IssuePolicyController } from './issue-policy/issue-policy.controller';
import { IssuePolicyHandler } from './issue-policy/issue-policy.handler';
import { RenewPolicyController } from './renew-policy/renew-policy.controller';
import { RenewPolicyHandler } from './renew-policy/renew-policy.handler';

// Añadir un caso de uso = añadir una carpeta y registrarla aquí.
@Module({
  controllers: [IssuePolicyController, RenewPolicyController, GetPolicyController],
  providers: [PolicyStore, NotificationService, IssuePolicyHandler, RenewPolicyHandler],
})
export class PolicyModule {}
