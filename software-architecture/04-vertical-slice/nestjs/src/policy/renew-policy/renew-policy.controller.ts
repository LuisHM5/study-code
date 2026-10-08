import { Body, Controller, HttpCode, Param, Post } from '@nestjs/common';
import { RenewPolicyDto } from './renew-policy.dto';
import { RenewPolicyHandler } from './renew-policy.handler';

@Controller('policies')
export class RenewPolicyController {
  constructor(private readonly handler: RenewPolicyHandler) {}

  @Post(':id/renew')
  @HttpCode(200)
  renew(@Param('id') id: string, @Body() dto: RenewPolicyDto) {
    return this.handler.execute(id, dto);
  }
}
