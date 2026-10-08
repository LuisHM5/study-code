import { Policy } from '../../domain/policy.entity';

export abstract class PolicyGateway {
  abstract existsByNumber(number: string): Promise<boolean>;
  abstract findById(id: string): Promise<Policy | null>;
  abstract save(policy: Policy): Promise<void>;
}
