import { Module } from '@nestjs/common';
import { HttpModule } from '@nestjs/axios';
import { ConfigModule } from '@nestjs/config';
import { TasksService } from './tasks.service';
import { TasksController } from './tasks.controller';

@Module({
  imports: [HttpModule, ConfigModule],
  controllers: [TasksController],
  providers: [TasksService],
})
export class TasksModule {}
