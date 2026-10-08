import { randomUUID } from 'node:crypto';
import { Injectable } from '@nestjs/common';
import { Policy } from '../domain/model/policy';
import { PolicyNumber } from '../domain/model/policy-number';
import { PolicyTerm } from '../domain/model/policy-term';
import { PolicyRepository } from '../domain/repositories/policy.repository';
import { PolicyUniquenessService } from '../domain/services/policy-uniqueness.service';
import { IssuePolicyCommand } from './dto/issue-policy.command';
import { PolicyNotifier } from './policy-notifier';

// Servicio de aplicación: orquesta, no decide reglas.
@Injectable()
export class IssuePolicyService {
  constructor(
    private readonly policies: PolicyRepository,
    private readonly uniqueness: PolicyUniquenessService,
    private readonly notifier: PolicyNotifier,
  ) {}

  async execute(cmd: IssuePolicyCommand): Promise<string> {
    const number = new PolicyNumber(cmd.number);
    await this.uniqueness.assertUnique(number);

    const policy = new Policy(randomUUID(), number, new PolicyTerm(cmd.startDate, cmd.endDate));
    await this.policies.save(policy);
    await this.notifier.policyIssued(policy);
    return policy.id;
  }
}
