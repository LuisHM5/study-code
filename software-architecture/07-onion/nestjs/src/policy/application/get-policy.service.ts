import { Injectable } from '@nestjs/common';
import { PolicyNotFoundError } from '../domain/errors';
import { Policy } from '../domain/model/policy';
import { PolicyRepository } from '../domain/repositories/policy.repository';

@Injectable()
export class GetPolicyService {
  constructor(private readonly policies: PolicyRepository) {}

  async execute(id: string): Promise<Policy> {
    const policy = await this.policies.findById(id);
    if (!policy) throw new PolicyNotFoundError();
    return policy;
  }
}
