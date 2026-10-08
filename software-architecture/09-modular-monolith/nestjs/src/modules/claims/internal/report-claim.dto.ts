import { IsNotEmpty, IsString } from 'class-validator';

export class ReportClaimDto {
  @IsString()
  @IsNotEmpty()
  policyId: string;

  @IsString()
  @IsNotEmpty()
  description: string;
}
