import { Body, Controller, Post } from '@nestjs/common';
import { IssuePolicyDto } from './issue-policy.dto';
import { IssuePolicyHandler } from './issue-policy.handler';

@Controller('policies')
export class IssuePolicyController {
  constructor(private readonly handler: IssuePolicyHandler) {}

  @Post()
  issue(@Body() dto: IssuePolicyDto) {
    return this.handler.execute(dto);
  }
}
