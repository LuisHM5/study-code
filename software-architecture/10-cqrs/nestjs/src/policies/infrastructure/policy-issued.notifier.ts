import { EventsHandler, IEventHandler } from '@nestjs/cqrs';
import { PolicyIssuedEvent } from '../domain/policy-issued.event';

@EventsHandler(PolicyIssuedEvent)
export class PolicyIssuedNotifier implements IEventHandler<PolicyIssuedEvent> {
  handle(event: PolicyIssuedEvent): void {
    console.log(`[notificación] póliza emitida ${event.number}`);
  }
}
