import { Body, Controller, Post } from '@nestjs/common';
import { ReportClaimDto } from './report-claim.dto';
import { ReportClaimUseCase } from './report-claim.use-case';

@Controller('claims')
export class ClaimController {
  constructor(private readonly reportClaim: ReportClaimUseCase) {}

  @Post()
  report(@Body() dto: ReportClaimDto) {
    return { id: this.reportClaim.execute(dto) };
  }
}
