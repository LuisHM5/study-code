import { ConflictException, Injectable } from '@nestjs/common';
import { PolicyEvent } from '../events/policy.events';

export type StoredEvent = PolicyEvent & { version: number; occurredAt: string };

// Event store en memoria: un stream (lista que solo crece) por agregado.
@Injectable()
export class EventStore {
  private readonly streams = new Map<string, StoredEvent[]>();

  read(streamId: string): StoredEvent[] {
    return this.streams.get(streamId) ?? [];
  }

  // Concurrencia optimista: solo se añade si nadie escribió desde que se leyó el stream.
  append(streamId: string, expectedVersion: number, events: PolicyEvent[]): StoredEvent[] {
    const stream = this.read(streamId);
    if (stream.length !== expectedVersion) {
      throw new ConflictException('El agregado fue modificado por otra operación');
    }
    const stored = events.map((event, i) => ({
      ...event,
      version: expectedVersion + i + 1,
      occurredAt: new Date().toISOString(),
    }));
    this.streams.set(streamId, [...stream, ...stored]);
    return stored;
  }
}
