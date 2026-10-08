import { PolicyNumber } from './policy-number';
import { PolicyTerm } from './policy-term';

// Centro de la cebolla: entidad con sus value objects. No depende de nada externo.
export class Policy {
  constructor(
    readonly id: string,
    readonly number: PolicyNumber,
    readonly term: PolicyTerm,
  ) {}
}
