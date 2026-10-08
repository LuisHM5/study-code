export interface Claim {
  id: string;
  policyId: string; // referencia por ID, sin relación directa a la entidad de policies
  description: string;
}
