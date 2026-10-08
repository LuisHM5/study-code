import { Injectable } from '@nestjs/common';
import { EventEmitter2 } from '@nestjs/event-emitter';
import { Policy } from '../domain/policy.aggregate';
import { EventStore } from './event-store';

@Injectable()
export class PolicyRepository {
  // Índice auxiliar para la regla de unicidad (el event store se lee por stream).
  private readonly numbers = new Set<string>();

  constructor(
    private readonly eventStore: EventStore,
    private readonly emitter: EventEmitter2,
  ) {}

  existsByNumber(number: string): boolean {
    return this.numbers.has(number);
  }

  // Cargar es reproducir los eventos del stream.
  load(id: string): Policy | null {
    const history = this.eventStore.read(`policy-${id}`);
    return history.length ? Policy.fromHistory(history) : null;
  }

  // Guardar es añadir los eventos nuevos y luego publicarlos a los listeners.
  save(policy: Policy): void {
    const changes = policy.pullChanges();
    const stored = this.eventStore.append(
      `policy-${policy.id}`,
      policy.version - changes.length,
      changes,
    );
    this.numbers.add(policy.number);
    for (const event of stored) this.emitter.emit(event.type, event);
  }
}
