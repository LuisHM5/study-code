import { Injectable } from '@nestjs/common';
import { Claim } from './claim';

@Injectable()
export class ClaimRepository {
  private readonly rows = new Map<string, Claim>();

  save(claim: Claim): void {
    this.rows.set(claim.id, claim);
  }
}
