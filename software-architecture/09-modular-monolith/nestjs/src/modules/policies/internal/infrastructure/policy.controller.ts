import { Body, Controller, Get, NotFoundException, Param, Post } from '@nestjs/common';
import { IssuePolicyUseCase } from '../application/issue-policy.use-case';
import { IssuePolicyDto } from './issue-policy.dto';
import { PolicyRepository } from './policy.repository';

@Controller('policies')
export class PolicyController {
  constructor(
    private readonly issuePolicy: IssuePolicyUseCase,
    private readonly policies: PolicyRepository,
  ) {}

  @Post()
  issue(@Body() dto: IssuePolicyDto) {
    const id = this.issuePolicy.execute({
      number: dto.number,
      startDate: new Date(dto.startDate),
      endDate: new Date(dto.endDate),
    });
    return { id };
  }

  @Get(':id')
  get(@Param('id') id: string) {
    const policy = this.policies.findById(id);
    if (!policy) throw new NotFoundException('Póliza no encontrada');
    return {
      id: policy.id,
      number: policy.number,
      startDate: policy.startDate.toISOString().slice(0, 10),
      endDate: policy.endDate.toISOString().slice(0, 10),
    };
  }
}
