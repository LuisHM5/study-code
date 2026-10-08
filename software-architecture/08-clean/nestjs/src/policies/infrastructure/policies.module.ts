import { Module } from '@nestjs/common';
import { APP_FILTER } from '@nestjs/core';
import { GetPolicyUseCase } from '../application/get-policy/get-policy.use-case';
import { IdGenerator } from '../application/gateways/id-generator';
import { NotificationGateway } from '../application/gateways/notification.gateway';
import { PolicyGateway } from '../application/gateways/policy.gateway';
import { IssuePolicyUseCase } from '../application/issue-policy/issue-policy.use-case';
import { DomainErrorFilter } from '../interface-adapters/controllers/domain-error.filter';
import { PolicyController } from '../interface-adapters/controllers/policy.controller';
import { ConsoleNotificationGateway } from '../interface-adapters/gateways/console-notification.gateway';
import { InMemoryPolicyGateway } from '../interface-adapters/gateways/in-memory-policy.gateway';
import { UuidGenerator } from '../interface-adapters/gateways/uuid-generator';

// Círculo 4: el único lugar que conoce todo y arma las piezas.
// Los casos de uso no tienen decoradores, por eso se construyen con fábricas.
@Module({
  controllers: [PolicyController],
  providers: [
    { provide: PolicyGateway, useClass: InMemoryPolicyGateway },
    { provide: NotificationGateway, useClass: ConsoleNotificationGateway },
    { provide: IdGenerator, useClass: UuidGenerator },
    {
      provide: IssuePolicyUseCase,
      useFactory: (p: PolicyGateway, n: NotificationGateway, i: IdGenerator) =>
        new IssuePolicyUseCase(p, n, i),
      inject: [PolicyGateway, NotificationGateway, IdGenerator],
    },
    {
      provide: GetPolicyUseCase,
      useFactory: (p: PolicyGateway) => new GetPolicyUseCase(p),
      inject: [PolicyGateway],
    },
    { provide: APP_FILTER, useClass: DomainErrorFilter },
  ],
})
export class PoliciesModule {}
