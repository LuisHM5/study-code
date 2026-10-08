import { Injectable } from '@nestjs/common';
import { Policy } from '../domain/policy';

// Lado de ESCRITURA
@Injectable()
export class PolicyRepository {
  private readonly rows = new Map<string, Policy>();

  existsByNumber(number: string): boolean {
    return [...this.rows.values()].some((p) => p.number === number);
  }

  save(policy: Policy): void {
    this.rows.set(policy.id, policy);
  }
}
