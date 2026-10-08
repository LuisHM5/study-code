import { randomUUID } from 'node:crypto';
import { ConflictException, Injectable } from '@nestjs/common';
import { EventEmitter2 } from '@nestjs/event-emitter';
import { POLICY_ISSUED, PolicyIssuedEvent } from '../../public/events/policy-issued.event';
import { Policy } from '../domain/policy';
import { PolicyRepository } from '../infrastructure/policy.repository';

@Injectable()
export class IssuePolicyUseCase {
  constructor(
    private readonly policies: PolicyRepository,
    private readonly events: EventEmitter2,
  ) {}

  execute(cmd: { number: string; startDate: Date; endDate: Date }): string {
    if (this.policies.existsByNumber(cmd.number)) throw new ConflictException('Número duplicado');

    const policy = Policy.issue(randomUUID(), cmd.number, cmd.startDate, cmd.endDate);
    this.policies.save(policy);
    // Publica el hecho; no sabe quién lo escucha (notifications).
    this.events.emit(POLICY_ISSUED, new PolicyIssuedEvent(policy.id, policy.number));
    return policy.id;
  }
}
