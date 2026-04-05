import { Injectable, HttpException, HttpStatus } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { ConfigService } from '@nestjs/config';
import { firstValueFrom } from 'rxjs';
import { AxiosError } from 'axios';
import { CreateTaskDto } from './dto/create-task.dto';
import { UpdateTaskDto } from './dto/update-task.dto';
import { GetTasksQueryDto } from './dto/get-tasks-query.dto';
import { Task } from './interfaces/task.interface';

@Injectable()
export class TasksService {
  private readonly baseUrl: string;

  constructor(
    private readonly httpService: HttpService,
    private readonly configService: ConfigService,
  ) {
    this.baseUrl = this.configService.get<string>('GO_SERVICE_URL') || 'http://localhost:8080';
  }

  async findAll(query: GetTasksQueryDto): Promise<Task[]> {
    try {
      const { data } = await firstValueFrom(
        this.httpService.get<Task[]>(`${this.baseUrl}/tasks`, { params: query }),
      );
      return data;
    } catch (error) {
      this.handleAxiosError(error);
    }
  }

  async create(createTaskDto: CreateTaskDto): Promise<Task> {
    try {
      const { data } = await firstValueFrom(
        this.httpService.post<Task>(`${this.baseUrl}/tasks`, createTaskDto),
      );
      return data;
    } catch (error) {
      this.handleAxiosError(error);
    }
  }

  async complete(id: string): Promise<Task> {
    try {
      const { data } = await firstValueFrom(
        this.httpService.patch<Task>(`${this.baseUrl}/tasks/${id}/complete`),
      );
      return data;
    } catch (error) {
      this.handleAxiosError(error);
    }
  }

  async remove(id: string): Promise<void> {
    try {
      await firstValueFrom(this.httpService.delete(`${this.baseUrl}/tasks/${id}`));
    } catch (error) {
      this.handleAxiosError(error);
    }
  }

  async update(id: string, updateTaskDto: UpdateTaskDto): Promise<Task> {
    try {
      const { data } = await firstValueFrom(
        this.httpService.put<Task>(`${this.baseUrl}/tasks/${id}`, updateTaskDto),
      );
      return data;
    } catch (error) {
      this.handleAxiosError(error);
    }
  }

  private handleAxiosError(error: any): never {
    if (error.isAxiosError) {
      const axiosError = error as AxiosError;
      const status = axiosError.response?.status || HttpStatus.INTERNAL_SERVER_ERROR;
      const message = axiosError.response?.data || axiosError.message;
      throw new HttpException(message, status);
    }
    throw new HttpException(
      'Internal Server Error',
      HttpStatus.INTERNAL_SERVER_ERROR,
    );
  }
}
