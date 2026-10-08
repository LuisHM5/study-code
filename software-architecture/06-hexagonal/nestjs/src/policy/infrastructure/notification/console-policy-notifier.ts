import { Injectable } from '@nestjs/common';
import { PolicyNotifier } from '../../application/ports/policy-notifier';
import { Policy } from '../../domain/policy';

// Adaptador de salida
@Injectable()
export class ConsolePolicyNotifier implements PolicyNotifier {
  async policyIssued(policy: Policy): Promise<void> {
    console.log(`[notificación] póliza emitida ${policy.number}`);
  }
}
