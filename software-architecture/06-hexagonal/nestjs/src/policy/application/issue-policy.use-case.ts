import { randomUUID } from 'node:crypto';
import { Injectable } from '@nestjs/common';
import { Policy } from '../domain/policy';
import { DuplicatedPolicyError } from '../domain/policy.errors';
import { PolicyNotifier } from './ports/policy-notifier';
import { PolicyRepository } from './ports/policy.repository';

export interface IssuePolicyCommand {
  number: string;
  startDate: Date;
  endDate: Date;
}

// Núcleo: solo conoce puertos, nunca adaptadores.
@Injectable()
export class IssuePolicyUseCase {
  constructor(
    private readonly policies: PolicyRepository,
    private readonly notifier: PolicyNotifier,
  ) {}

  async execute(cmd: IssuePolicyCommand): Promise<string> {
    if (await this.policies.existsByNumber(cmd.number)) throw new DuplicatedPolicyError();

    const policy = Policy.issue(randomUUID(), cmd.number, cmd.startDate, cmd.endDate);
    await this.policies.save(policy);
    await this.notifier.policyIssued(policy);
    return policy.id;
  }
}
