import { BadRequestException } from '@nestjs/common';

// Regla pura compartida por los slices: se prueba sin mocks.
export function assertValidTerm(start: Date, end: Date): void {
  if (end <= start) throw new BadRequestException('Vigencia inválida');
}

export function addMonths(date: string, months: number): string {
  const result = new Date(date);
  result.setUTCMonth(result.getUTCMonth() + months);
  return result.toISOString().slice(0, 10);
}
