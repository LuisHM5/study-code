// Lo único importable desde fuera del módulo. Todo lo de internal/ es privado.
export { PoliciesModule } from './policies.module';
export { PoliciesFacade } from './public/policies.facade';
export { PolicySummaryDto } from './public/policy.dto';
export { POLICY_ISSUED, PolicyIssuedEvent } from './public/events/policy-issued.event';
