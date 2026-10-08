import { Module } from '@nestjs/common';
import { PoliciesModule } from '../policies';
import { ClaimController } from './internal/claim.controller';
import { ClaimRepository } from './internal/claim.repository';
import { ReportClaimUseCase } from './internal/report-claim.use-case';

@Module({
  imports: [PoliciesModule],
  controllers: [ClaimController],
  providers: [ReportClaimUseCase, ClaimRepository],
})
export class ClaimsModule {}
