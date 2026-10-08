import {
  Body,
  Controller,
  DefaultValuePipe,
  Get,
  Param,
  ParseIntPipe,
  Post,
  Query,
} from '@nestjs/common';
import { CommandBus, QueryBus } from '@nestjs/cqrs';
import { IssuePolicyCommand } from './commands/issue-policy/issue-policy.command';
import { IssuePolicyDto } from './issue-policy.dto';
import { GetPolicyQuery } from './queries/get-policy/get-policy.query';
import { ListExpiringPoliciesQuery } from './queries/list-expiring-policies/list-expiring-policies.query';

// El controlador solo despacha: comandos por un bus, consultas por otro.
@Controller('policies')
export class PolicyController {
  constructor(
    private readonly commandBus: CommandBus,
    private readonly queryBus: QueryBus,
  ) {}

  @Post()
  async issue(@Body() dto: IssuePolicyDto) {
    const id = await this.commandBus.execute(
      new IssuePolicyCommand(dto.number, new Date(dto.startDate), new Date(dto.endDate)),
    );
    return { id };
  }

  // Debe declararse antes de ':id' para que "expiring" no se tome como un ID.
  @Get('expiring')
  expiring(@Query('days', new DefaultValuePipe(30), ParseIntPipe) days: number) {
    return this.queryBus.execute(new ListExpiringPoliciesQuery(days));
  }

  @Get(':id')
  get(@Param('id') id: string) {
    return this.queryBus.execute(new GetPolicyQuery(id));
  }
}
