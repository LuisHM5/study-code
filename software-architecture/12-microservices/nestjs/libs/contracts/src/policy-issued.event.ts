// Contrato compartido entre servicios. Cambiarlo exige coordinar a ambos lados.
export interface PolicyIssuedEvent {
  policyId: string;
  number: string;
}
