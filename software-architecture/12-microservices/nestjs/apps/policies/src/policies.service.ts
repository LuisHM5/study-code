import { randomUUID } from 'node:crypto';
import {
  BadRequestException,
  ConflictException,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import { IssuePolicyDto } from './issue-policy.dto';
import { NotificationsClient } from './notifications.client';

export interface Policy {
  id: string;
  number: string;
  startDate: string;
  endDate: string;
}

@Injectable()
export class PoliciesService {
  // Cada servicio es dueño de su propia base de datos (aquí, en memoria).
  private readonly rows = new Map<string, Policy>();

  constructor(private readonly notifications: NotificationsClient) {}

  async issue(dto: IssuePolicyDto): Promise<Policy> {
    if (new Date(dto.endDate) <= new Date(dto.startDate)) {
      throw new BadRequestException('Vigencia inválida');
    }
    if ([...this.rows.values()].some((p) => p.number === dto.number)) {
      throw new ConflictException('Número duplicado');
    }

    const policy: Policy = { id: randomUUID(), ...dto };
    this.rows.set(policy.id, policy);
    await this.notifications.policyIssued({ policyId: policy.id, number: policy.number });
    return policy;
  }

  getById(id: string): Policy {
    const policy = this.rows.get(id);
    if (!policy) throw new NotFoundException('Póliza no encontrada');
    return policy;
  }
}
