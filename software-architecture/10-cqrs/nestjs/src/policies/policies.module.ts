import { Module } from '@nestjs/common';
import { CqrsModule } from '@nestjs/cqrs';
import { IssuePolicyHandler } from './commands/issue-policy/issue-policy.handler';
import { PolicyIssuedNotifier } from './infrastructure/policy-issued.notifier';
import { PolicyReadModel } from './infrastructure/policy-read-model';
import { PolicyReadModelProjector } from './infrastructure/policy-read-model.projector';
import { PolicyRepository } from './infrastructure/policy.repository';
import { PolicyController } from './policy.controller';
import { GetPolicyHandler } from './queries/get-policy/get-policy.handler';
import { ListExpiringPoliciesHandler } from './queries/list-expiring-policies/list-expiring-policies.handler';

@Module({
  imports: [CqrsModule],
  controllers: [PolicyController],
  providers: [
    // Escritura
    PolicyRepository,
    IssuePolicyHandler,
    // Lectura
    PolicyReadModel,
    PolicyReadModelProjector,
    GetPolicyHandler,
    ListExpiringPoliciesHandler,
    // Efectos secundarios
    PolicyIssuedNotifier,
  ],
})
export class PoliciesModule {}
