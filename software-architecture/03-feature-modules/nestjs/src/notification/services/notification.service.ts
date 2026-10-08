import { Injectable } from '@nestjs/common';

@Injectable()
export class NotificationService {
  policyIssued(policyNumber: string): void {
    console.log(`[notificación] póliza emitida ${policyNumber}`);
  }
}
