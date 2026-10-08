import { Injectable } from '@nestjs/common';
import { Policy } from '../domain/policy';
import { PolicyNotFoundError } from '../domain/policy.errors';
import { PolicyRepository } from './ports/policy.repository';

@Injectable()
export class GetPolicyUseCase {
  constructor(private readonly policies: PolicyRepository) {}

  async execute(id: string): Promise<Policy> {
    const policy = await this.policies.findById(id);
    if (!policy) throw new PolicyNotFoundError();
    return policy;
  }
}
