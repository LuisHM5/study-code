import { DuplicatedPolicyError } from '../errors';
import { PolicyNumber } from '../model/policy-number';
import { PolicyRepository } from '../repositories/policy.repository';

// Servicio de dominio: regla que una entidad sola no puede verificar.
// Sin @Injectable(): el dominio no conoce NestJS (se registra con una fábrica en el módulo).
export class PolicyUniquenessService {
  constructor(private readonly policies: PolicyRepository) {}

  async assertUnique(number: PolicyNumber): Promise<void> {
    if (await this.policies.existsByNumber(number)) throw new DuplicatedPolicyError();
  }
}
