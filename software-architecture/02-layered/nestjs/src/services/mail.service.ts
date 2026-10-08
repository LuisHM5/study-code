import { Injectable } from '@nestjs/common';
import { Policy } from '../entities/policy.entity';

@Injectable()
export class MailService {
  policyIssued(policy: Policy): void {
    console.log(`[notificación] póliza emitida ${policy.number}`);
  }
}
