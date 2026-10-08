import { DomainError } from '../errors';

// Value object
export class PolicyNumber {
  readonly value: string;

  constructor(value: string) {
    const normalized = value.trim().toUpperCase();
    if (!normalized) throw new DomainError('El número de póliza es obligatorio');
    this.value = normalized;
  }

  equals(other: PolicyNumber): boolean {
    return this.value === other.value;
  }
}
