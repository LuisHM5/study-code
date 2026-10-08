import { Module } from '@nestjs/common';
import { APP_FILTER } from '@nestjs/core';
import { GetPolicyService } from './application/get-policy.service';
import { IssuePolicyService } from './application/issue-policy.service';
import { PolicyNotifier } from './application/policy-notifier';
import { PolicyRepository } from './domain/repositories/policy.repository';
import { PolicyUniquenessService } from './domain/services/policy-uniqueness.service';
import { ConsolePolicyNotifier } from './infrastructure/notification/console-policy-notifier';
import { InMemoryPolicyRepository } from './infrastructure/persistence/in-memory-policy.repository';
import { DomainErrorFilter } from './presentation/domain-error.filter';
import { PolicyController } from './presentation/policy.controller';

@Module({
  controllers: [PolicyController],
  providers: [
    IssuePolicyService,
    GetPolicyService,
    { provide: PolicyRepository, useClass: InMemoryPolicyRepository },
    { provide: PolicyNotifier, useClass: ConsolePolicyNotifier },
    // El servicio de dominio no tiene decoradores: se construye con una fábrica.
    {
      provide: PolicyUniquenessService,
      useFactory: (repo: PolicyRepository) => new PolicyUniquenessService(repo),
      inject: [PolicyRepository],
    },
    { provide: APP_FILTER, useClass: DomainErrorFilter },
  ],
})
export class PolicyModule {}
