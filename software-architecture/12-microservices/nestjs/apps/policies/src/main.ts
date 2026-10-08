import 'reflect-metadata';
import { ValidationPipe } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { PoliciesModule } from './policies.module';

async function bootstrap() {
  const app = await NestFactory.create(PoliciesModule);
  app.useGlobalPipes(new ValidationPipe({ whitelist: true }));
  await app.listen(process.env.PORT ?? 3001);
}
bootstrap();
