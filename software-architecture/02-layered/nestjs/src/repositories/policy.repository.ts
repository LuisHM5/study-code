import { Injectable } from '@nestjs/common';
import { Policy } from '../entities/policy.entity';

// Capa de datos. En memoria; en un proyecto real sería el repositorio del ORM.
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
