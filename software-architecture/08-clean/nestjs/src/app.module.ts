import { Module } from '@nestjs/common';
import { PoliciesModule } from './policies/infrastructure/policies.module';

@Module({ imports: [PoliciesModule] })
export class AppModule {}
