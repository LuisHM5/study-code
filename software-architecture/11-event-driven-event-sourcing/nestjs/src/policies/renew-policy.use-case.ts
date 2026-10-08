import { Injectable, NotFoundException } from '@nestjs/common';
import { Policy } from './domain/policy.aggregate';
import { PolicyRepository } from './infrastructure/policy.repository';

@Injectable()
export class RenewPolicyUseCase {
  constructor(private readonly policies: PolicyRepository) {}

  execute(id: string, months: number): Policy {
    const policy = this.policies.load(id);
    if (!policy) throw new NotFoundException('Póliza no encontrada');

    policy.renew(months);
    this.policies.save(policy);
    return policy;
  }
}
