import { Controller, Get, NotFoundException, Param } from '@nestjs/common';
import { PolicyStore } from '../entities/policy.store';

// Un slice simple no necesita handler: consulta directo.
@Controller('policies')
export class GetPolicyController {
  constructor(private readonly store: PolicyStore) {}

  @Get(':id')
  get(@Param('id') id: string) {
    const policy = this.store.findById(id);
    if (!policy) throw new NotFoundException('Póliza no encontrada');
    return policy;
  }
}
