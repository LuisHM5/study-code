import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { GetPolicyUseCase } from '../../application/get-policy/get-policy.use-case';
import { IssuePolicyUseCase } from '../../application/issue-policy/issue-policy.use-case';
import { PolicyPresenter } from '../presenters/policy.presenter';
import { IssuePolicyRequest } from './issue-policy.request';

@Controller('policies')
export class PolicyController {
  constructor(
    private readonly issuePolicy: IssuePolicyUseCase,
    private readonly getPolicy: GetPolicyUseCase,
  ) {}

  @Post()
  async issue(@Body() req: IssuePolicyRequest) {
    // Request HTTP -> Input del caso de uso
    const output = await this.issuePolicy.execute({
      number: req.number,
      startDate: new Date(req.startDate),
      endDate: new Date(req.endDate),
    });
    return PolicyPresenter.issued(output);
  }

  @Get(':id')
  async get(@Param('id') id: string) {
    return PolicyPresenter.detail(await this.getPolicy.execute(id));
  }
}
