import { randomUUID } from 'node:crypto';
import { ConflictException } from '@nestjs/common';
import { CommandHandler, EventBus, ICommandHandler } from '@nestjs/cqrs';
import { Policy } from '../../domain/policy';
import { PolicyIssuedEvent } from '../../domain/policy-issued.event';
import { PolicyRepository } from '../../infrastructure/policy.repository';
import { IssuePolicyCommand } from './issue-policy.command';

// Comando: cambia el estado pasando por las reglas del dominio. Devuelve solo el ID.
@CommandHandler(IssuePolicyCommand)
export class IssuePolicyHandler implements ICommandHandler<IssuePolicyCommand, string> {
  constructor(
    private readonly policies: PolicyRepository,
    private readonly eventBus: EventBus,
  ) {}

  async execute(cmd: IssuePolicyCommand): Promise<string> {
    if (this.policies.existsByNumber(cmd.number)) throw new ConflictException('Número duplicado');

    const policy = Policy.issue(randomUUID(), cmd.number, cmd.startDate, cmd.endDate);
    this.policies.save(policy);
    this.eventBus.publish(
      new PolicyIssuedEvent(policy.id, policy.number, policy.startDate, policy.endDate),
    );
    return policy.id;
  }
}
