export class PolicyIssuedEvent {
  constructor(
    readonly policyId: string,
    readonly number: string,
    readonly startDate: Date,
    readonly endDate: Date,
  ) {}
}
