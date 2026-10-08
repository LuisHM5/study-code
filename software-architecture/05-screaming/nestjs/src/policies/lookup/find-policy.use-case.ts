import { Injectable, NotFoundException } from '@nestjs/common';
import { Policy } from '../policy.entity';
import { PolicyStore } from '../policy.store';

@Injectable()
export class FindPolicy {
  constructor(private readonly policies: PolicyStore) {}

  execute(id: string): Policy {
    const policy = this.policies.findById(id);
    if (!policy) throw new NotFoundException('Póliza no encontrada');
    return policy;
  }
}
