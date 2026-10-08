// Errores de negocio: TypeScript puro, sin HTTP ni NestJS.
export class DomainError extends Error {}
export class InvalidTermError extends DomainError {
  constructor() {
    super('La vigencia es inválida');
  }
}
export class DuplicatedPolicyError extends DomainError {
  constructor() {
    super('El número de póliza ya existe');
  }
}
export class PolicyNotFoundError extends DomainError {
  constructor() {
    super('Póliza no encontrada');
  }
}
