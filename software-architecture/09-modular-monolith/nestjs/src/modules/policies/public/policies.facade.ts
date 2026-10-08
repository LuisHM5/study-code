import { Injectable } from '@nestjs/common';
import { PolicyRepository } from '../internal/infrastructure/policy.repository';
import { PolicySummaryDto } from './policy.dto';

// API pública del módulo: lo único que otros módulos pueden usar.
@Injectable()
export class PoliciesFacade {
  constructor(private readonly policies: PolicyRepository) {}

  findActiveById(id: string): PolicySummaryDto | null {
    const policy = this.policies.findById(id);
    if (!policy || !policy.isActiveOn(new Date())) return null;
    return { id: policy.id, number: policy.number };
  }
}
