import { BadRequestException } from '@nestjs/common';
import { PolicyEvent } from '../events/policy.events';

// Agregado con event sourcing: el estado NO se guarda, se reconstruye aplicando eventos.
export class Policy {
  id: string;
  number: string;
  startDate: string;
  endDate: string;
  version = 0; // cuántos eventos se han aplicado

  private changes: PolicyEvent[] = [];

  // Reconstruye el estado actual a partir del historial.
  static fromHistory(history: PolicyEvent[]): Policy {
    const policy = new Policy();
    for (const event of history) policy.apply(event);
    return policy;
  }

  // Comandos: validan y registran un evento; no tocan el estado directamente.
  static issue(id: string, number: string, startDate: string, endDate: string): Policy {
    if (new Date(endDate) <= new Date(startDate)) {
      throw new BadRequestException('Vigencia inválida');
    }
    const policy = new Policy();
    policy.record({ type: 'policy.issued', data: { policyId: id, number, startDate, endDate } });
    return policy;
  }

  renew(months: number): void {
    const end = new Date(this.endDate);
    end.setUTCMonth(end.getUTCMonth() + months);
    this.record({
      type: 'policy.renewed',
      data: { policyId: this.id, months, newEndDate: end.toISOString().slice(0, 10) },
    });
  }

  // Eventos nuevos pendientes de guardar.
  pullChanges(): PolicyEvent[] {
    const changes = this.changes;
    this.changes = [];
    return changes;
  }

  private record(event: PolicyEvent): void {
    this.apply(event);
    this.changes.push(event);
  }

  // Único lugar donde cambia el estado, tanto al ejecutar como al reproducir.
  private apply(event: PolicyEvent): void {
    switch (event.type) {
      case 'policy.issued':
        this.id = event.data.policyId;
        this.number = event.data.number;
        this.startDate = event.data.startDate;
        this.endDate = event.data.endDate;
        break;
      case 'policy.renewed':
        this.endDate = event.data.newEndDate;
        break;
    }
    this.version++;
  }
}
