import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { GetPolicyService } from '../application/get-policy.service';
import { IssuePolicyService } from '../application/issue-policy.service';
import { IssuePolicyRequest } from './issue-policy.request';

@Controller('policies')
export class PolicyController {
  constructor(
    private readonly issuePolicy: IssuePolicyService,
    private readonly getPolicy: GetPolicyService,
  ) {}

  @Post()
  async issue(@Body() req: IssuePolicyRequest) {
    const id = await this.issuePolicy.execute({
      number: req.number,
      startDate: new Date(req.startDate),
      endDate: new Date(req.endDate),
    });
    return { id };
  }

  @Get(':id')
  async get(@Param('id') id: string) {
    const policy = await this.getPolicy.execute(id);
    return {
      id: policy.id,
      number: policy.number.value,
      startDate: policy.term.start.toISOString().slice(0, 10),
      endDate: policy.term.end.toISOString().slice(0, 10),
    };
  }
}
