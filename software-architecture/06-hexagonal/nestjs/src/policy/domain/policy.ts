import { InvalidTermError } from './policy.errors';

// Modelo de dominio: sin decoradores ni dependencias externas.
export class Policy {
  private constructor(
    readonly id: string,
    readonly number: string,
    readonly startDate: Date,
    readonly endDate: Date,
  ) {}

  static issue(id: string, number: string, startDate: Date, endDate: Date): Policy {
    if (endDate <= startDate) throw new InvalidTermError();
    return new Policy(id, number, startDate, endDate);
  }
}
