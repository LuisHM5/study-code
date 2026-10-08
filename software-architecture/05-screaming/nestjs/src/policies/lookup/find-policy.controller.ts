import { Controller, Get, Param } from '@nestjs/common';
import { FindPolicy } from './find-policy.use-case';

@Controller('policies')
export class FindPolicyController {
  constructor(private readonly findPolicy: FindPolicy) {}

  @Get(':id')
  get(@Param('id') id: string) {
    return this.findPolicy.execute(id);
  }
}
