import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { IssuePolicyDto } from '../dtos/issue-policy.dto';
import { PolicyService } from '../services/policy.service';

// Capa HTTP: sin lógica, solo traduce la petición a una llamada de servicio.
@Controller('policies')
export class PolicyController {
  constructor(private readonly policyService: PolicyService) {}

  @Post()
  issue(@Body() dto: IssuePolicyDto) {
    const policy = this.policyService.issue(dto);
    return { id: policy.id };
  }

  @Get(':id')
  get(@Param('id') id: string) {
    return this.policyService.getById(id);
  }
}
