import { IsDateString, IsNotEmpty, IsString } from 'class-validator';

export class IssuePolicyDto {
  @IsString()
  @IsNotEmpty()
  number: string;

  @IsDateString()
  startDate: string;

  @IsDateString()
  endDate: string;
}
