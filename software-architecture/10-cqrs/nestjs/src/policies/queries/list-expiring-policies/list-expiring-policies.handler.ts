import { IQueryHandler, QueryHandler } from '@nestjs/cqrs';
import { PolicyReadModel } from '../../infrastructure/policy-read-model';
import { ExpiringPolicyView } from './expiring-policy.view';
import { ListExpiringPoliciesQuery } from './list-expiring-policies.query';

const DAY_MS = 24 * 60 * 60 * 1000;

@QueryHandler(ListExpiringPoliciesQuery)
export class ListExpiringPoliciesHandler
  implements IQueryHandler<ListExpiringPoliciesQuery, ExpiringPolicyView[]>
{
  constructor(private readonly readModel: PolicyReadModel) {}

  async execute(query: ListExpiringPoliciesQuery): Promise<ExpiringPolicyView[]> {
    const today = new Date().toISOString().slice(0, 10);
    const limit = new Date(Date.parse(today) + query.days * DAY_MS).toISOString().slice(0, 10);

    return this.readModel.expiringBetween(today, limit).map((v) => ({
      id: v.id,
      number: v.number,
      endDate: v.endDate,
      daysLeft: Math.round((Date.parse(v.endDate) - Date.parse(today)) / DAY_MS),
    }));
  }
}
