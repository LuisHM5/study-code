import { PolicyNotFoundError } from '../../domain/policy.errors';
import { PolicyGateway } from '../gateways/policy.gateway';
import { GetPolicyOutput } from './get-policy.output';

export class GetPolicyUseCase {
  constructor(private readonly policies: PolicyGateway) {}

  async execute(id: string): Promise<GetPolicyOutput> {
    const policy = await this.policies.findById(id);
    if (!policy) throw new PolicyNotFoundError();
    // La entidad no cruza la frontera: sale una estructura simple.
    return {
      id: policy.id,
      number: policy.number,
      startDate: policy.startDate,
      endDate: policy.endDate,
    };
  }
}
