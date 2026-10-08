// Modelo de entrada del caso de uso: no es el DTO HTTP ni la entidad.
export interface IssuePolicyInput {
  number: string;
  startDate: Date;
  endDate: Date;
}
