import { randomUUID } from 'node:crypto';
import {
  BadRequestException,
  ConflictException,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import { NotificationService } from '../../notification/services/notification.service';
import { IssuePolicyDto } from '../dto/issue-policy.dto';
import { Policy } from '../entities/policy.entity';
import { PolicyRepository } from '../repositories/policy.repository';

@Injectable()
export class PolicyService {
  constructor(
    private readonly repo: PolicyRepository,
    // Viene de otro módulo: solo se puede inyectar porque NotificationModule lo exporta.
    private readonly notifications: NotificationService,
  ) {}

  issue(dto: IssuePolicyDto): Policy {
    if (new Date(dto.endDate) <= new Date(dto.startDate)) {
      throw new BadRequestException('Vigencia inválida');
    }
    if (this.repo.existsByNumber(dto.number)) {
      throw new ConflictException('Número duplicado');
    }

    const policy = this.repo.save({ id: randomUUID(), ...dto });
    this.notifications.policyIssued(policy.number);
    return policy;
  }

  getById(id: string): Policy {
    const policy = this.repo.findById(id);
    if (!policy) throw new NotFoundException('Póliza no encontrada');
    return policy;
  }
}
