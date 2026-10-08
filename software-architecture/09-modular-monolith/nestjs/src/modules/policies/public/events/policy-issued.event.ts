export const POLICY_ISSUED = 'policy.issued';

export class PolicyIssuedEvent {
  constructor(
    readonly policyId: string,
    readonly number: string,
  ) {}
}
