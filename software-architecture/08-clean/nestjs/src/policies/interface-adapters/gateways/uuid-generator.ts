import { randomUUID } from 'node:crypto';
import { IdGenerator } from '../../application/gateways/id-generator';

export class UuidGenerator implements IdGenerator {
  next(): string {
    return randomUUID();
  }
}
