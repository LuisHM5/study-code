import { InvalidTermError } from '../errors';

// Value object: inmutable y siempre válido.
export class PolicyTerm {
  constructor(
    readonly start: Date,
    readonly end: Date,
  ) {
    if (end <= start) throw new InvalidTermError();
  }

  isActiveOn(date: Date): boolean {
    return date >= this.start && date <= this.end;
  }
}
