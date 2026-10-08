import 'reflect-metadata';
import { NestFactory } from '@nestjs/core';
import { NotificationsModule } from './notifications.module';

// Servicio independiente: su propio proceso, puerto y ciclo de despliegue.
async function bootstrap() {
  const app = await NestFactory.create(NotificationsModule);
  await app.listen(process.env.PORT ?? 3002);
}
bootstrap();
