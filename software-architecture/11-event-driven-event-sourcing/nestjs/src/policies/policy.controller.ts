import { Body, Controller, Get, HttpCode, NotFoundException, Param, Post } from '@nestjs/common';
import { EventStore } from './infrastructure/event-store';
import { PolicyRepository } from './infrastructure/policy.repository';
import { IssuePolicyDto } from './issue-policy.dto';
import { IssuePolicyUseCase } from './issue-policy.use-case';
import { RenewPolicyDto } from './renew-policy.dto';
import { RenewPolicyUseCase } from './renew-policy.use-case';

@Controller('policies')
export class PolicyController {
  constructor(
    private readonly issuePolicy: IssuePolicyUseCase,
    private readonly renewPolicy: RenewPolicyUseCase,
    private readonly policies: PolicyRepository,
    private readonly eventStore: EventStore,
  ) {}

  @Post()
  issue(@Body() dto: IssuePolicyDto) {
    return { id: this.issuePolicy.execute(dto) };
  }

  @Post(':id/renew')
  @HttpCode(200)
  renew(@Param('id') id: string, @Body() dto: RenewPolicyDto) {
    const policy = this.renewPolicy.execute(id, dto.months);
    return { id: policy.id, endDate: policy.endDate };
  }

  // El estado se reconstruye desde los eventos en cada lectura.
  @Get(':id')
  get(@Param('id') id: string) {
    const policy = this.policies.load(id);
    if (!policy) throw new NotFoundException('Póliza no encontrada');
    return {
      id: policy.id,
      number: policy.number,
      startDate: policy.startDate,
      endDate: policy.endDate,
    };
  }

  // El historial completo: la auditoría sale gratis.
  @Get(':id/events')
  events(@Param('id') id: string) {
    const events = this.eventStore.read(`policy-${id}`);
    if (!events.length) throw new NotFoundException('Póliza no encontrada');
    return events;
  }
}
