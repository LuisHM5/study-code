import { Policy } from '../../domain/policy';

// Puerto de salida
export abstract class PolicyNotifier {
  abstract policyIssued(policy: Policy): Promise<void>;
}
