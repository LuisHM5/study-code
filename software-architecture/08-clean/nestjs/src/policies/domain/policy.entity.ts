import { InvalidTermError } from './policy.errors';

// Círculo 1: reglas de negocio. No importa nada de fuera de domain/.
export class Policy {
  constructor(
    readonly id: string,
    readonly number: string,
    readonly startDate: Date,
    readonly endDate: Date,
  ) {
    if (endDate <= startDate) throw new InvalidTermError();
  }
}
