import { randomUUID } from 'node:crypto';
import {
  BadRequestException,
  ConflictException,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import { IssuePolicyDto } from '../dtos/issue-policy.dto';
import { Policy } from '../entities/policy.entity';
import { PolicyRepository } from '../repositories/policy.repository';
import { MailService } from './mail.service';

// Capa de lógica: no conoce HTTP, solo llama a la capa de abajo.
@Injectable()
export class PolicyService {
  constructor(
    private readonly repo: PolicyRepository,
    private readonly mailService: MailService,
  ) {}

  issue(dto: IssuePolicyDto): Policy {
    if (new Date(dto.endDate) <= new Date(dto.startDate)) {
      throw new BadRequestException('Vigencia inválida');
    }
    if (this.repo.existsByNumber(dto.number)) {
      throw new ConflictException('Número duplicado');
    }

    const policy = this.repo.save({ id: randomUUID(), ...dto });
    this.mailService.policyIssued(policy);
    return policy;
  }

  getById(id: string): Policy {
    const policy = this.repo.findById(id);
    if (!policy) throw new NotFoundException('Póliza no encontrada');
    return policy;
  }
}
