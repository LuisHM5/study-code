import { Policy } from '../../domain/policy';

// Puerto de salida. Clase abstracta: sirve de interfaz y de token de inyección.
export abstract class PolicyRepository {
  abstract existsByNumber(number: string): Promise<boolean>;
  abstract findById(id: string): Promise<Policy | null>;
  abstract save(policy: Policy): Promise<void>;
}
