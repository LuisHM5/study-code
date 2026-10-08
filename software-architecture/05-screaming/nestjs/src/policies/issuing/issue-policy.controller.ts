import { Body, Controller, Post } from '@nestjs/common';
import { IssuePolicyDto } from './issue-policy.dto';
import { IssuePolicy } from './issue-policy.use-case';

@Controller('policies')
export class IssuePolicyController {
  constructor(private readonly issuePolicy: IssuePolicy) {}

  @Post()
  issue(@Body() dto: IssuePolicyDto) {
    return { id: this.issuePolicy.execute(dto).id };
  }
}
