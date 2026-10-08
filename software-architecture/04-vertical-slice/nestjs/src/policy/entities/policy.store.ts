import { Injectable } from '@nestjs/common';
import { Policy } from './policy.entity';

// Almacén en memoria compartido por los slices (hace el papel del Repository<PolicyEntity>).
@Injectable()
export class PolicyStore {
  private readonly rows = new Map<string, Policy>();

  existsByNumber(number: string): boolean {
    return [...this.rows.values()].some((p) => p.number === number);
  }

  findById(id: string): Policy | undefined {
    return this.rows.get(id);
  }

  save(policy: Policy): Policy {
    this.rows.set(policy.id, policy);
    return policy;
  }
}
