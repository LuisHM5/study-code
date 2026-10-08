import { Injectable } from '@nestjs/common';
import { Policy } from '../../domain/model/policy';
import { PolicyNumber } from '../../domain/model/policy-number';
import { PolicyRepository } from '../../domain/repositories/policy.repository';

@Injectable()
export class InMemoryPolicyRepository implements PolicyRepository {
  private readonly rows = new Map<string, Policy>();

  async existsByNumber(number: PolicyNumber): Promise<boolean> {
    return [...this.rows.values()].some((p) => p.number.equals(number));
  }

  async findById(id: string): Promise<Policy | null> {
    return this.rows.get(id) ?? null;
  }

  async save(policy: Policy): Promise<void> {
    this.rows.set(policy.id, policy);
  }
}
