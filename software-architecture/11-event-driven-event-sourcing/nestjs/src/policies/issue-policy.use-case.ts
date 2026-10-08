import { randomUUID } from 'node:crypto';
import { ConflictException, Injectable } from '@nestjs/common';
import { Policy } from './domain/policy.aggregate';
import { PolicyRepository } from './infrastructure/policy.repository';
import { IssuePolicyDto } from './issue-policy.dto';

@Injectable()
export class IssuePolicyUseCase {
  constructor(private readonly policies: PolicyRepository) {}

  execute(dto: IssuePolicyDto): string {
    if (this.policies.existsByNumber(dto.number)) throw new ConflictException('Número duplicado');

    const policy = Policy.issue(randomUUID(), dto.number, dto.startDate, dto.endDate);
    this.policies.save(policy); // persiste y publica 'policy.issued'
    return policy.id;
  }
}
