import { Injectable, NotFoundException } from '@nestjs/common';
import { addMonths } from '../domain/policy-term';
import { PolicyStore } from '../entities/policy.store';
import { RenewPolicyDto } from './renew-policy.dto';

@Injectable()
export class RenewPolicyHandler {
  constructor(private readonly store: PolicyStore) {}

  execute(id: string, dto: RenewPolicyDto): { id: string; endDate: string } {
    const policy = this.store.findById(id);
    if (!policy) throw new NotFoundException('Póliza no encontrada');

    policy.endDate = addMonths(policy.endDate, dto.months);
    this.store.save(policy);
    return { id: policy.id, endDate: policy.endDate };
  }
}
