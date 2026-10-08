import { Module } from '@nestjs/common';
import { PoliciesModule } from './policies/policies.module';

// El árbol de src/ dice de qué trata el negocio: policies/issuing, policies/lookup.
@Module({ imports: [PoliciesModule] })
export class AppModule {}
