import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { IssuePolicyDto } from './issue-policy.dto';
import { PoliciesService } from './policies.service';

@Controller('policies')
export class PoliciesController {
  constructor(private readonly policies: PoliciesService) {}

  @Post()
  async issue(@Body() dto: IssuePolicyDto) {
    const policy = await this.policies.issue(dto);
    return { id: policy.id };
  }

  @Get(':id')
  get(@Param('id') id: string) {
    return this.policies.getById(id);
  }
}
