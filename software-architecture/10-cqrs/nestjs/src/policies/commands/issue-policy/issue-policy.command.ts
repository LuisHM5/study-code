export class IssuePolicyCommand {
  constructor(
    readonly number: string,
    readonly startDate: Date,
    readonly endDate: Date,
  ) {}
}
