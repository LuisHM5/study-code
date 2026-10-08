import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { GetPolicyUseCase } from '../../application/get-policy.use-case';
import { IssuePolicyUseCase } from '../../application/issue-policy.use-case';
import { IssuePolicyDto } from './issue-policy.dto';

// Adaptador de entrada: traduce HTTP a una llamada al núcleo.
@Controller('policies')
export class PolicyController {
  constructor(
    private readonly issuePolicy: IssuePolicyUseCase,
    private readonly getPolicy: GetPolicyUseCase,
  ) {}

  @Post()
  async issue(@Body() dto: IssuePolicyDto) {
    const id = await this.issuePolicy.execute({
      number: dto.number,
      startDate: new Date(dto.startDate),
      endDate: new Date(dto.endDate),
    });
    return { id };
  }

  @Get(':id')
  async get(@Param('id') id: string) {
    const policy = await this.getPolicy.execute(id);
    return {
      id: policy.id,
      number: policy.number,
      startDate: policy.startDate.toISOString().slice(0, 10),
      endDate: policy.endDate.toISOString().slice(0, 10),
    };
  }
}
