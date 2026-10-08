import { BadRequestException } from '@nestjs/common';

// Modelo de escritura: solo lo necesario para cumplir las reglas.
export class Policy {
  private constructor(
    readonly id: string,
    readonly number: string,
    readonly startDate: Date,
    readonly endDate: Date,
  ) {}

  static issue(id: string, number: string, startDate: Date, endDate: Date): Policy {
    if (endDate <= startDate) throw new BadRequestException('Vigencia inválida');
    return new Policy(id, number, startDate, endDate);
  }
}
