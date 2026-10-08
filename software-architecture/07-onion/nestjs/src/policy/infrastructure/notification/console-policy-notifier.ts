import { Injectable } from '@nestjs/common';
import { PolicyNotifier } from '../../application/policy-notifier';
import { Policy } from '../../domain/model/policy';

@Injectable()
export class ConsolePolicyNotifier implements PolicyNotifier {
  async policyIssued(policy: Policy): Promise<void> {
    console.log(`[notificación] póliza emitida ${policy.number.value}`);
  }
}
