import { randomUUID } from 'node:crypto';
import { Injectable, UnprocessableEntityException } from '@nestjs/common';
// Se importa desde el index del otro módulo, nunca desde su carpeta internal/.
import { PoliciesFacade } from '../../policies';
import { ClaimRepository } from './claim.repository';
import { ReportClaimDto } from './report-claim.dto';

@Injectable()
export class ReportClaimUseCase {
  constructor(
    private readonly policies: PoliciesFacade,
    private readonly claims: ClaimRepository,
  ) {}

  execute(dto: ReportClaimDto): string {
    const policy = this.policies.findActiveById(dto.policyId);
    if (!policy) throw new UnprocessableEntityException('La póliza no existe o no está vigente');

    const id = randomUUID();
    this.claims.save({ id, policyId: policy.id, description: dto.description });
    return id;
  }
}
