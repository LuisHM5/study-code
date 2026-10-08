import { Injectable } from '@nestjs/common';
import { Policy } from '../entities/policy.entity';

// Privado del módulo policy (no se exporta). En memoria en lugar de TypeORM.
@Injectable()
export class PolicyRepository {
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
