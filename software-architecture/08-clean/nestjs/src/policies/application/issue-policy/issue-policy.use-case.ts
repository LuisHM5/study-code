import { Policy } from '../../domain/policy.entity';
import { DuplicatedPolicyError } from '../../domain/policy.errors';
import { IdGenerator } from '../gateways/id-generator';
import { NotificationGateway } from '../gateways/notification.gateway';
import { PolicyGateway } from '../gateways/policy.gateway';
import { IssuePolicyInput } from './issue-policy.input';
import { IssuePolicyOutput } from './issue-policy.output';

// Círculo 2: sin imports de NestJS ni de infraestructura. Solo domain/ y sus gateways.
export class IssuePolicyUseCase {
  constructor(
    private readonly policies: PolicyGateway,
    private readonly notifications: NotificationGateway,
    private readonly ids: IdGenerator,
  ) {}

  async execute(input: IssuePolicyInput): Promise<IssuePolicyOutput> {
    if (await this.policies.existsByNumber(input.number)) throw new DuplicatedPolicyError();

    const policy = new Policy(this.ids.next(), input.number, input.startDate, input.endDate);
    await this.policies.save(policy);
    await this.notifications.policyIssued(policy.number);
    return { id: policy.id, number: policy.number };
  }
}
