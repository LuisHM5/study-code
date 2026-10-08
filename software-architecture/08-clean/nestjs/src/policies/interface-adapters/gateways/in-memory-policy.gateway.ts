import { PolicyGateway } from '../../application/gateways/policy.gateway';
import { Policy } from '../../domain/policy.entity';

// Círculo 3: implementa el gateway que define el caso de uso.
export class InMemoryPolicyGateway implements PolicyGateway {
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
