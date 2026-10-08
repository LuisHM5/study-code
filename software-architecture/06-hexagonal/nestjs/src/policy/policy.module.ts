import { Module } from '@nestjs/common';
import { APP_FILTER } from '@nestjs/core';
import { GetPolicyUseCase } from './application/get-policy.use-case';
import { IssuePolicyUseCase } from './application/issue-policy.use-case';
import { PolicyNotifier } from './application/ports/policy-notifier';
import { PolicyRepository } from './application/ports/policy.repository';
import { DomainErrorFilter } from './infrastructure/http/domain-error.filter';
import { PolicyController } from './infrastructure/http/policy.controller';
import { ConsolePolicyNotifier } from './infrastructure/notification/console-policy-notifier';
import { InMemoryPolicyRepository } from './infrastructure/persistence/in-memory-policy.repository';

@Module({
  controllers: [PolicyController],
  providers: [
    IssuePolicyUseCase,
    GetPolicyUseCase,
    // Aquí se conecta cada puerto con su adaptador
    { provide: PolicyRepository, useClass: InMemoryPolicyRepository },
    { provide: PolicyNotifier, useClass: ConsolePolicyNotifier },
    { provide: APP_FILTER, useClass: DomainErrorFilter },
  ],
})
export class PolicyModule {}
