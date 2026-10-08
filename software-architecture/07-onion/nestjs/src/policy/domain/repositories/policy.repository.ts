import { Policy } from '../model/policy';
import { PolicyNumber } from '../model/policy-number';

// La interfaz vive en el dominio; la implementación, en el anillo exterior.
export abstract class PolicyRepository {
  abstract existsByNumber(number: PolicyNumber): Promise<boolean>;
  abstract findById(id: string): Promise<Policy | null>;
  abstract save(policy: Policy): Promise<void>;
}
