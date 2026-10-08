import { IsInt, Min } from 'class-validator';

export class RenewPolicyDto {
  @IsInt()
  @Min(1)
  months: number;
}
