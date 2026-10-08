import { Module } from '@nestjs/common';
import { IssuePolicyUseCase } from './internal/application/issue-policy.use-case';
import { PolicyController } from './internal/infrastructure/policy.controller';
import { PolicyRepository } from './internal/infrastructure/policy.repository';
import { PoliciesFacade } from './public/policies.facade';

@Module({
  controllers: [PolicyController],
  providers: [PoliciesFacade, IssuePolicyUseCase, PolicyRepository],
  exports: [PoliciesFacade], // solo la fachada sale del módulo
})
export class PoliciesModule {}
