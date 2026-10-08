import { Injectable, Logger } from '@nestjs/common';
import { PolicyIssuedEvent } from '../../../libs/contracts/src/policy-issued.event';

// Lo que antes era una llamada a función ahora es una llamada de red:
// puede tardar o fallar, así que lleva timeout y no tumba la emisión.
@Injectable()
export class NotificationsClient {
  private readonly logger = new Logger(NotificationsClient.name);
  private readonly baseUrl = process.env.NOTIFICATIONS_URL ?? 'http://localhost:3002';

  async policyIssued(event: PolicyIssuedEvent): Promise<void> {
    try {
      const res = await fetch(`${this.baseUrl}/notifications`, {
        method: 'POST',
        headers: { 'content-type': 'application/json' },
        body: JSON.stringify(event),
        signal: AbortSignal.timeout(2000),
      });
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
    } catch (error) {
      this.logger.warn(`No se pudo notificar la póliza ${event.number}: ${error}`);
    }
  }
}
