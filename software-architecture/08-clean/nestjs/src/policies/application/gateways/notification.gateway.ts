export abstract class NotificationGateway {
  abstract policyIssued(policyNumber: string): Promise<void>;
}
