import { Injectable } from '@nestjs/common';
import { OnEvent } from '@nestjs/event-emitter';
import { PolicyIssued } from '../../policies/events/policy.events';

// Event-driven: reacciona al hecho. El módulo de pólizas no sabe que este listener existe.
@Injectable()
export class PolicyIssuedListener {
  @OnEvent('policy.issued')
  handle(event: PolicyIssued): void {
    console.log(`[notificación] póliza emitida ${event.data.number}`);
  }
}
