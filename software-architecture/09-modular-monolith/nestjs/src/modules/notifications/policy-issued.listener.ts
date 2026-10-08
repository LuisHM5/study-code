import { Injectable } from '@nestjs/common';
import { OnEvent } from '@nestjs/event-emitter';
import { POLICY_ISSUED, PolicyIssuedEvent } from '../policies';

// Reacciona al evento sin que policies sepa que este módulo existe.
@Injectable()
export class PolicyIssuedListener {
  @OnEvent(POLICY_ISSUED)
  handle(event: PolicyIssuedEvent): void {
    console.log(`[notificación] póliza emitida ${event.number}`);
  }
}
