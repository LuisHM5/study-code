import { GetPolicyOutput } from '../../application/get-policy/get-policy.output';
import { IssuePolicyOutput } from '../../application/issue-policy/issue-policy.output';

// Convierte la salida del caso de uso al formato que espera el cliente HTTP.
export class PolicyPresenter {
  static issued(output: IssuePolicyOutput) {
    return { id: output.id };
  }

  static detail(output: GetPolicyOutput) {
    return {
      id: output.id,
      number: output.number,
      startDate: output.startDate.toISOString().slice(0, 10),
      endDate: output.endDate.toISOString().slice(0, 10),
    };
  }
}
