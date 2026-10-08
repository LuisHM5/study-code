import { randomUUID } from 'node:crypto';
import { BadRequestException, ConflictException, Injectable } from '@nestjs/common';
import { NotificationService } from '../../shared/notifications/notification.service';
import { Policy } from '../policy.entity';
import { PolicyStore } from '../policy.store';
import { IssuePolicyDto } from './issue-policy.dto';

// La clase se llama como la acción de negocio: "emitir póliza".
@Injectable()
export class IssuePolicy {
  constructor(
    private readonly policies: PolicyStore,
    private readonly notifications: NotificationService,
  ) {}

  execute(dto: IssuePolicyDto): Policy {
    if (new Date(dto.endDate) <= new Date(dto.startDate)) {
      throw new BadRequestException('Vigencia inválida');
    }
    if (this.policies.existsByNumber(dto.number)) {
      throw new ConflictException('Número duplicado');
    }

    const policy = this.policies.save({ id: randomUUID(), ...dto });
    this.notifications.policyIssued(policy.number);
    return policy;
  }
}
