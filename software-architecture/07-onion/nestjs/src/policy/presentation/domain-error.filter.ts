import { ArgumentsHost, Catch, ExceptionFilter, HttpStatus } from '@nestjs/common';
import {
  DomainError,
  DuplicatedPolicyError,
  PolicyNotFoundError,
} from '../domain/errors';

// Traduce errores de dominio a HTTP. Cualquier otro DomainError es un dato inválido (400).
@Catch(DomainError)
export class DomainErrorFilter implements ExceptionFilter {
  catch(error: DomainError, host: ArgumentsHost) {
    let status = HttpStatus.BAD_REQUEST;
    if (error instanceof DuplicatedPolicyError) status = HttpStatus.CONFLICT;
    if (error instanceof PolicyNotFoundError) status = HttpStatus.NOT_FOUND;

    host.switchToHttp().getResponse().status(status).json({ message: error.message });
  }
}
