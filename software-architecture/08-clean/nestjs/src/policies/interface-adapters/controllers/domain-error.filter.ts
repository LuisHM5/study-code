import { ArgumentsHost, Catch, ExceptionFilter, HttpStatus } from '@nestjs/common';
import {
  DomainError,
  DuplicatedPolicyError,
  InvalidTermError,
  PolicyNotFoundError,
} from '../../domain/policy.errors';

// Traduce errores de dominio a HTTP. El dominio no sabe de códigos de estado.
@Catch(DomainError)
export class DomainErrorFilter implements ExceptionFilter {
  catch(error: DomainError, host: ArgumentsHost) {
    let status = HttpStatus.INTERNAL_SERVER_ERROR;
    if (error instanceof InvalidTermError) status = HttpStatus.BAD_REQUEST;
    if (error instanceof DuplicatedPolicyError) status = HttpStatus.CONFLICT;
    if (error instanceof PolicyNotFoundError) status = HttpStatus.NOT_FOUND;

    host.switchToHttp().getResponse().status(status).json({ message: error.message });
  }
}
