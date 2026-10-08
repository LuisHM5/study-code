import { NotificationGateway } from '../../application/gateways/notification.gateway';

export class ConsoleNotificationGateway implements NotificationGateway {
  async policyIssued(policyNumber: string): Promise<void> {
    console.log(`[notificación] póliza emitida ${policyNumber}`);
  }
}
