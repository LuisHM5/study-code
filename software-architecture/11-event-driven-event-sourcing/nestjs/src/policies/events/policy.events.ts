// Eventos de dominio: hechos en pasado, inmutables. Son lo único que se persiste.
export interface PolicyIssued {
  type: 'policy.issued';
  data: { policyId: string; number: string; startDate: string; endDate: string };
}

export interface PolicyRenewed {
  type: 'policy.renewed';
  data: { policyId: string; months: number; newEndDate: string };
}

export type PolicyEvent = PolicyIssued | PolicyRenewed;
