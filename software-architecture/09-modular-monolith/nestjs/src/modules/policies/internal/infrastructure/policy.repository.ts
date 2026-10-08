import { Injectable } from '@nestjs/common';
import { Policy } from '../domain/policy';

// El módulo es dueño de sus datos: nadie más toca este almacén.
@Injectable()
export class PolicyRepository {
  private readonly rows = new Map<string, Policy>();

  existsByNumber(number: string): boolean {
    return [...this.rows.values()].some((p) => p.number === number);
  }

  findById(id: string): Policy | undefined {
    return this.rows.get(id);
  }

  save(policy: Policy): void {
    this.rows.set(policy.id, policy);
  }
}
