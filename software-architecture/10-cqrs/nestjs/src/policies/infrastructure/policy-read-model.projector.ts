import { EventsHandler, IEventHandler } from '@nestjs/cqrs';
import { PolicyIssuedEvent } from '../domain/policy-issued.event';
import { PolicyReadModel } from './policy-read-model';

// Proyección: mantiene el modelo de lectura al día a partir de los eventos.
@EventsHandler(PolicyIssuedEvent)
export class PolicyReadModelProjector implements IEventHandler<PolicyIssuedEvent> {
  constructor(private readonly readModel: PolicyReadModel) {}

  handle(event: PolicyIssuedEvent): void {
    this.readModel.upsert({
      id: event.policyId,
      number: event.number,
      startDate: event.startDate.toISOString().slice(0, 10),
      endDate: event.endDate.toISOString().slice(0, 10),
    });
  }
}
