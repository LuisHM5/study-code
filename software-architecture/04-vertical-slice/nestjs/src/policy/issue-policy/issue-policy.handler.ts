import { randomUUID } from 'node:crypto';
import { ConflictException, Injectable } from '@nestjs/common';
import { NotificationService } from '../../notification/notification.service';
import { assertValidTerm } from '../domain/policy-term';
import { PolicyStore } from '../entities/policy.store';
import { IssuePolicyDto } from './issue-policy.dto';

// Un caso de uso = una clase. Solo recibe las dependencias que este slice usa.
@Injectable()
export class IssuePolicyHandler {
  constructor(
    private readonly store: PolicyStore,
    private readonly notifications: NotificationService,
  ) {}

  execute(dto: IssuePolicyDto): { id: string } {
    assertValidTerm(new Date(dto.startDate), new Date(dto.endDate));
    if (this.store.existsByNumber(dto.number)) {
      throw new ConflictException('Número duplicado');
    }

    const policy = this.store.save({ id: randomUUID(), ...dto });
    this.notifications.policyIssued(policy.number);
    return { id: policy.id };
  }
}
