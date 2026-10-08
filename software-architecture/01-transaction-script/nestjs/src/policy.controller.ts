import { randomUUID } from 'node:crypto';
import {
  BadRequestException,
  Body,
  ConflictException,
  Controller,
  Get,
  NotFoundException,
  Param,
  Post,
} from '@nestjs/common';
import { IssuePolicyDto } from './issue-policy.dto';
import { Policy } from './policy.entity';

// Transaction Script: el handler hace todo (validar, guardar, notificar).
@Controller('policies')
export class PolicyController {
  // "Base de datos" en memoria
  private readonly policies = new Map<string, Policy>();

  @Post()
  issue(@Body() dto: IssuePolicyDto) {
    if (new Date(dto.endDate) <= new Date(dto.startDate)) {
      throw new BadRequestException('La vigencia es inválida');
    }
    for (const existing of this.policies.values()) {
      if (existing.number === dto.number) {
        throw new ConflictException('El número de póliza ya existe');
      }
    }

    const policy: Policy = { id: randomUUID(), ...dto };
    this.policies.set(policy.id, policy);
    console.log(`[notificación] póliza emitida ${policy.number}`);
    return { id: policy.id };
  }

  @Get(':id')
  get(@Param('id') id: string) {
    const policy = this.policies.get(id);
    if (!policy) throw new NotFoundException('Póliza no encontrada');
    return policy;
  }
}
