import { Module } from '@nestjs/common';
import { NotificationModule } from '../notification/notification.module';
import { PolicyController } from './controllers/policy.controller';
import { PolicyRepository } from './repositories/policy.repository';
import { PolicyService } from './services/policy.service';

// Todo lo de "pólizas" vive en esta carpeta y se declara en este módulo.
@Module({
  imports: [NotificationModule],
  controllers: [PolicyController],
  providers: [PolicyService, PolicyRepository],
  exports: [PolicyService], // la API pública del módulo
})
export class PolicyModule {}
