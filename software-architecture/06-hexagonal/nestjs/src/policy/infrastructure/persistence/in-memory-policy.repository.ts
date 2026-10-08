import { Injectable } from '@nestjs/common';
import { PolicyRepository } from '../../application/ports/policy.repository';
import { Policy } from '../../domain/policy';

// Adaptador de salida. Cambiarlo por uno de TypeORM no toca el núcleo.
@Injectable()
export class InMemoryPolicyRepository implements PolicyRepository {
  private readonly rows = new Map<string, Policy>();

  async existsByNumber(number: string): Promise<boolean> {
    return [...this.rows.values()].some((p) => p.number === number);
  }

  async findById(id: string): Promise<Policy | null> {
    return this.rows.get(id) ?? null;
  }

  async save(policy: Policy): Promise<void> {
    this.rows.set(policy.id, policy);
  }
}
