import { Policy } from '../domain/model/policy';

// Interfaz que la aplicación necesita del anillo exterior.
export abstract class PolicyNotifier {
  abstract policyIssued(policy: Policy): Promise<void>;
}
