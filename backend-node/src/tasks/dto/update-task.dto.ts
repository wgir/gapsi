import { IsOptional, IsString, IsEnum, MinLength } from 'class-validator';

export enum TaskStatus {
  TODO = 'TODO',
  DONE = 'DONE',
  CANCELLED = 'CANCELLED',
}

export class UpdateTaskDto {
  @IsString()
  @IsOptional()
  @MinLength(3)
  title?: string;

  @IsString()
  @IsOptional()
  description?: string;

  @IsEnum(TaskStatus)
  @IsOptional()
  status?: TaskStatus;
}
