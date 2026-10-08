import { Module } from '@nestjs/common';
import { EventStore } from './infrastructure/event-store';
import { PolicyRepository } from './infrastructure/policy.repository';
import { IssuePolicyUseCase } from './issue-policy.use-case';
import { PolicyController } from './policy.controller';
import { RenewPolicyUseCase } from './renew-policy.use-case';

@Module({
  controllers: [PolicyController],
  providers: [EventStore, PolicyRepository, IssuePolicyUseCase, RenewPolicyUseCase],
})
export class PoliciesModule {}
