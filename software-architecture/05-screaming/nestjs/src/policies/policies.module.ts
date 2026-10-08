import { Module } from '@nestjs/common';
import { NotificationService } from '../shared/notifications/notification.service';
import { IssuePolicyController } from './issuing/issue-policy.controller';
import { IssuePolicy } from './issuing/issue-policy.use-case';
import { FindPolicyController } from './lookup/find-policy.controller';
import { FindPolicy } from './lookup/find-policy.use-case';
import { PolicyStore } from './policy.store';

@Module({
  controllers: [IssuePolicyController, FindPolicyController],
  providers: [PolicyStore, NotificationService, IssuePolicy, FindPolicy],
})
export class PoliciesModule {}
