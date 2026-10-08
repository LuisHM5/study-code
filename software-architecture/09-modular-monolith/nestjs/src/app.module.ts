import { Module } from '@nestjs/common';
import { EventEmitterModule } from '@nestjs/event-emitter';
import { ClaimsModule } from './modules/claims';
import { NotificationsModule } from './modules/notifications';
import { PoliciesModule } from './modules/policies';

// Un solo proceso, tres módulos con fronteras: API pública (fachada) y eventos.
@Module({
  imports: [EventEmitterModule.forRoot(), PoliciesModule, ClaimsModule, NotificationsModule],
})
export class AppModule {}
