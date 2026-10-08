import { NotFoundException } from '@nestjs/common';
import { IQueryHandler, QueryHandler } from '@nestjs/cqrs';
import { PolicyReadModel, PolicyView } from '../../infrastructure/policy-read-model';
import { GetPolicyQuery } from './get-policy.query';

// Consulta: lee directo del modelo de lectura, sin pasar por el dominio.
@QueryHandler(GetPolicyQuery)
export class GetPolicyHandler implements IQueryHandler<GetPolicyQuery, PolicyView> {
  constructor(private readonly readModel: PolicyReadModel) {}

  async execute(query: GetPolicyQuery): Promise<PolicyView> {
    const view = this.readModel.findById(query.id);
    if (!view) throw new NotFoundException('Póliza no encontrada');
    return view;
  }
}
